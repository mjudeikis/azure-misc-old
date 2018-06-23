.PHONY: all

all: glide build-purge

glide: 
	glide install -v

build-purge: 
	cd cmd/azure-purge ; go build -o main .

build-purge-local:
	docker build -f Dockerfile.azure-purge -t azure-purge .