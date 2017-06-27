all: vendor build

.PHONY: build
build:
	go build

.PHONY: vendor
vendor:
	glide install