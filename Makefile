GOPATH:=$(shell go env GOPATH)
GOOS:=$(shell go env GOOS)
GOARCH:=$(shell go env GOARCH)

all: build

.PHONY: build
build: deps
	CGO_ENABLED=0 go build -o sops-getJinkensTags -v

.PHONY: test
test:
	go test -v ./... -cover

deps:
	go get  github.com/golang/glog
	go get  github.com/micro/go-micro
	go get  github.com/golang/glog
	go get  github.com/micro/go-micro/metadata
	go get  github.com/micro/go-grpc
	go get  github.com/spf13/viper


clean:
	go clean -x gitlab.hfjy.com/infr/sops-getJinkensTags/sops-getJinkensTags

.PHONY: docker
docker: build
	docker build . -t sops-getJinkensTags:latest
