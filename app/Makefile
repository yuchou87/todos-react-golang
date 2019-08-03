.PHONY: all
all: build-go build-images push-images

GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/app/release
APPNAME=app

REPOSITORY=localhost:32000
APP_IMAGE=$(REPOSITORY)/app:latest

build-go: clean
	@echo " > Building binary..."
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=on go build -mod vendor -v -a -o $(GOBIN)/$(GOOS)/$(GOARCH)/$(APPNAME)

build-images: build-go
	@echo " > Building docker images... "
	@docker build -t $(APP_IMAGE) .

push-images:
	@docker push $(APP_IMAGE)

clean:
	@echo " > Cleaning build cache..."
	@GO111MODULE=on go clean
	@-rm -rf $(GOBIN)