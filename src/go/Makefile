.PHONY: all

all: glide build-purge

glide: 
	glide install -v

build-purge: 
	cd cmd/azure-purge ; go build .

build-purge-local:
	docker build -f Dockerfile.azure-purge -t azure-purge .