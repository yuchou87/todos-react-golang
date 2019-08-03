.PHONY: all
all: build-go build-images push-images

build-go: clean
	@echo " > Building binary..."
	@$(MAKE) -C app


build-web:
	@echo "> Building web..."
	@$(MAKE) -C web

deploy-images:
	@echo " > Deploy docker images... "