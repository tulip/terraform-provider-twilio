NAME:=$(shell basename $$PWD)
ARCH:=$(shell uname -m)
REPO:=$(shell git config --get remote.origin.url | perl -ne 'm{github.com[:/](.+/[^.]+)}; print  $$1')
VERSION=0.0.2

build:
	go build

build_linux:
	docker-machine start || true
	bash -c 'eval $$(docker-machine env) ; rocker build --no-garbage'

release: build build_linux
	rm -rf release && mkdir release
	mkdir -p build/Darwin && mv $(NAME) build/Darwin
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(ARCH).tgz -C build/Darwin $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(ARCH).tgz -C build/Linux $(NAME)
	gh-release create $(REPO) $(VERSION) master

.PHONY: release build build_linux
