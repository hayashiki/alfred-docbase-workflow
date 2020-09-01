VERSION := $$(make -s show-version)
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
ARCHIVE := "DocBase-$(VERSION).alfredworkflow"

GOBIN ?= $(shell go env GOPATH)/bin
export GO111MODULE ?= on

.PHONY: show-version
show-version: $(GOBIN)/gobump
	@gobump show -r .

.PHONY: build
build:
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o "./.workflow/dc"; \
	VERSION=v$(VERSION) envsubst >./.workflow/info.plist <./.workflow/info.plist.template;
#	sed -e 's/$(VERSION)/$(VERSION)/' info.plist
	zip -r "./bin/$(ARCHIVE)" ./.workflow/*;
	zip -d "./bin/$(ARCHIVE)" ./.workflow/info.plist.template;

.PHONY: tag
tag:
	git tag -a "v$(VERSION)" -m "Release $(VERSION)"
	git push --tags

