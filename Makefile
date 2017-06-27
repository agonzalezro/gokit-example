all: vendor build

.PHONY: build
build: build-server build-cli

.PHONY: build-cli
build-cli:
	go build ./cmd/cli

.PHONY: build-server
build-server:
	go build ./cmd/server

.PHONY: proto
proto:
	cd pb&&protoc whatever.proto --go_out=plugins=grpc:.

.PHONY: vendor
vendor:
	glide install