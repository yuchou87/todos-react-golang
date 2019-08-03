.PHONY: all
all: build-go build-images push-images

GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/app/release

REPOSITORY=localhost:32000
APP_IMAGE=$(REPOSITORY)/app:latest
WEB_IMAGE=$(REPOSITORY)/web:latest

build-go: clean
	@echo " > Building binary..."
	@cd app; CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) GO111MODULE=on go build -mod vendor -v -a -o $(GOBIN)/$(GOOS)/$(GOARCH)/$(APPNAME); cd -

build-web:
	@echo "> Building web..."
	@cd web; yarn build;cd -

build-images: build-go
	@echo " > Building docker images... "
	@git pull
	@cd app; docker build -t $(APP_IMAGE) .; cd -
	@cd web; docker build -t $(WEB_IMAGE) .; cd -

push-images:
	@docker push $(APP_IMAGE)
	@docker push $(WEB_IMAGE)

deploy-images:
	@echo " > Deploy docker images... "

clean:
	@echo " > Cleaning build cache..."
	@GO111MODULE=on go clean
	@-rm -rf $(GOBIN)