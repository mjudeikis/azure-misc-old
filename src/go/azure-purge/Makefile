.PHONY: all

all: glide build

glide: 
	glide install -v

build: 
	go build -o main .

build_local:
	docker build -t azure-purge .