.PHONY: all
all: build-web build-images push-images

REPOSITORY=localhost:32000
WEB_IMAGE=$(REPOSITORY)/web:latest

build-web:
	@echo "> Building web..."
	@yarn build

build-images:
	@echo " > Building docker images... "
	@docker build -t $(WEB_IMAGE) .

push-images:
	@docker push $(WEB_IMAGE)

clean:
	@echo " > Cleaning build cache..."
	@-rm -rf build