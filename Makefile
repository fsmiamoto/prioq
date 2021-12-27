GO := $(shell which go1.18beta1)

prelude:
	@$(GO) version
	@date

build: prelude
	$(GO) build

test: prelude
	$(GO) test -v ./...

fmt: prelude
	$(GO) fmt .

ci: fmt test build

.PHONY: prelude test fmt build ci
