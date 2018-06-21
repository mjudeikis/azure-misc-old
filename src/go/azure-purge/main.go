package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-04-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-02-01/resources"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2017-10-01/storage"
	azstorage "github.com/Azure/azure-sdk-for-go/storage"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const (
	resourceGroup  = "images"
	storageAccount = "openshiftimages"
	container      = "images"
	keepImages     = 5
	buildTimeout   = 6 * time.Hour
	groupTimeout   = 3 * 24 * time.Hour
)

var dryRun = flag.Bool("n", false, "dry-run")

type byName []compute.Image

func (b byName) Len() int           { return len(b) }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byName) Less(i, j int) bool { return *b[i].Name < *b[j].Name }

var clients = struct {
	accounts storage.AccountsClient
	groups   resources.GroupsClient
	images   compute.ImagesClient
	storage  azstorage.Client
}{}

var now = time.Now()

func getClients() error {
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")

	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return err
	}

	clients.accounts = storage.NewAccountsClient(subscriptionID)
	clients.accounts.Authorizer = authorizer
	clients.groups = resources.NewGroupsClient(subscriptionID)
	clients.groups.Authorizer = authorizer
	clients.images = compute.NewImagesClient(subscriptionID)
	clients.images.Authorizer = authorizer

	keys, err := clients.accounts.ListKeys(context.Background(), resourceGroup, storageAccount)
	if err != nil {
		return err
	}

	clients.storage, err = azstorage.NewClient(storageAccount, *(*keys.Keys)[0].Value, azstorage.DefaultBaseURL, azstorage.DefaultAPIVersion, true)
	if err != nil {
		return err
	}

	return nil
}

func listGroups() ([]resources.Group, error) {
	results, err := clients.groups.List(context.Background(), "", nil)
	if err != nil {
		return nil, err
	}

	var groups []resources.Group
	for ; results.NotDone(); results.Next() {
		groups = append(groups, results.Values()...)
	}

	return groups, nil
}

func listImages() ([]compute.Image, error) {
	results, err := clients.images.ListByResourceGroup(context.Background(), resourceGroup)
	if err != nil {
		return nil, err
	}

	var images []compute.Image
	for ; results.NotDone(); results.Next() {
		images = append(images, results.Values()...)
	}

	return images, nil
}

func deleteGroups(groups []resources.Group) error {
	var futures []resources.GroupsDeleteFuture
	for _, group := range groups {
		fmt.Printf("delete group %s\n", *group.Name)
		if *dryRun {
			continue
		}

		future, err := clients.groups.Delete(context.Background(), *group.Name)
		if err != nil {
			return err
		}
		futures = append(futures, future)
	}

	for _, future := range futures {
		err := future.WaitForCompletion(context.Background(), clients.groups.Client)
		if err != nil {
			return err
		}
	}

	return nil
}

func deleteImages(images []compute.Image) error {
	var futures []compute.ImagesDeleteFuture
	for _, image := range images {
		fmt.Printf("delete image %s\n", *image.Name)
		if *dryRun {
			continue
		}

		future, err := clients.images.Delete(context.Background(), resourceGroup, *image.Name)
		if err != nil {
			return err
		}
		futures = append(futures, future)
	}

	for _, future := range futures {
		err := future.WaitForCompletion(context.Background(), clients.images.Client)
		if err != nil {
			return err
		}
	}

	return nil
}

// purgeInvalidImages removes images from the "images" resourcegroup that are
// not tagged "valid: true" and which are older than `buildTimeout`.
func purgeInvalidImages() error {
	imageRx := regexp.MustCompile(`^.*-([0-9]{12})$`)

	images, err := listImages()
	if err != nil {
		return err
	}

	var toDelete []compute.Image
	for _, image := range images {
		m := imageRx.FindStringSubmatch(*image.Name)
		if m == nil {
			toDelete = append(toDelete, image)
			continue
		}

		t, err := time.Parse("200601021504", m[1])
		if err == nil && now.Sub(t) < buildTimeout {
			continue
		}

		v := image.Tags["valid"]
		if v == nil || *v != "true" {
			toDelete = append(toDelete, image)
		}
	}

	return deleteImages(toDelete)
}

// purgeOldImages removes images from the "images" resourcegroup, leaving only
// the `keepImages` most recent images of each kind.
func purgeOldImages() error {
	imageRx := regexp.MustCompile(`^(.*)-[0-9]{12}$`)

	images, err := listImages()
	if err != nil {
		return err
	}

	sort.Sort(sort.Reverse(byName(images)))

	var toDelete []compute.Image
	var lastPrefix *string
	var i int
	for _, image := range images {
		m := imageRx.FindStringSubmatch(*image.Name)
		switch {
		case m == nil:
			toDelete = append(toDelete, image)
		case lastPrefix == nil || m[1] != *lastPrefix:
			lastPrefix = &m[1]
			i = 1
		default:
			i++
			if i > keepImages {
				toDelete = append(toDelete, image)
			}
		}
	}

	return deleteImages(toDelete)
}

// purgeBlobs removes all blobs from `storageAccount`/`container` which do not
// have a matching image in `resourceGroup` and which are older than
// `buildTimeout`.
func purgeBlobs() error {
	blobRx := regexp.MustCompile(`-([0-9]{12})\.vhd$`)

	images, err := listImages()
	if err != nil {
		return err
	}
	allowedBlobs := make(map[string]struct{}, len(images))
	for _, image := range images {
		allowedBlobs[*image.Name+".vhd"] = struct{}{}
	}

	bs := clients.storage.GetBlobService()
	ctr := bs.GetContainerReference(container)

	blobs, err := ctr.ListBlobs(azstorage.ListBlobsParameters{})
	if err != nil {
		return err
	}

	for _, blob := range blobs.Blobs {
		if _, allowed := allowedBlobs[blob.Name]; allowed {
			continue
		}
		if m := blobRx.FindStringSubmatch(blob.Name); m != nil {
			t, err := time.Parse("200601021504", m[1])
			if err == nil && now.Sub(t) < buildTimeout {
				continue
			}
		}
		fmt.Printf("delete blob %s\n", blob.Name)
		if *dryRun {
			continue
		}

		if err = blob.Delete(nil); err != nil {
			return err
		}
	}

	return nil
}

// purgeGroups removes all resource groups tagged with the "now" tag, where the
// tag time is older than `groupTimeout`.
func purgeGroups() error {
	groups, err := listGroups()
	if err != nil {
		return err
	}

	var toDelete []resources.Group
	for _, group := range groups {
		timestamp := group.Tags["now"]
		if timestamp == nil {
			continue
		}
		t, err := strconv.ParseInt(*timestamp, 10, 64)
		if err == nil && now.Sub(time.Unix(t, 0)) < groupTimeout {
			continue
		}
		toDelete = append(toDelete, group)
	}

	return deleteGroups(toDelete)
}

func run() error {
	if err := getClients(); err != nil {
		return err
	}

	if err := purgeInvalidImages(); err != nil {
		return err
	}

	if err := purgeOldImages(); err != nil {
		return err
	}

	if err := purgeBlobs(); err != nil {
		return err
	}

	if err := purgeGroups(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Start azure-purge")
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}
