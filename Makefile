# Should be at least Go 1.16
GO := $(shell which go)

build: prelude
	$(GO) build

test: prelude
	$(GO) test -v ./...

fmt: prelude
	$(GO) fmt .

prelude:
	@$(GO) version
	@date

ci: fmt test build

.PHONY: prelude test fmt build ci
