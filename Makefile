VERSION := $$(make -s show-version)
CURRENT_REVISION := $(shell git rev-parse --short HEAD)
ARCHIVE := "DocBase-v$(VERSION).alfredworkflow"

GOBIN ?= $(shell go env GOPATH)/bin
export GO111MODULE ?= on

.PHONY: show-version
show-version: $(GOBIN)/gobump
	@gobump show -r .

$(GOBIN)/gobump:
	@cd && go get github.com/x-motemen/gobump/cmd/gobump

.PHONY: build
build:
	GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w" -o "./.workflow/dc"; \
	VERSION=v$(VERSION) envsubst >./.workflow/info.plist <./.workflow/info.plist.template;
	zip -r "./bin/$(ARCHIVE)" ./.workflow/*;
	zip -d "./bin/$(ARCHIVE)" ./.workflow/info.plist.template;

.PHONY: tag
tag:
	git tag -a "v$(VERSION)" -m "Release $(VERSION)"
	git push --tags

