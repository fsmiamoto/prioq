# Should be at least Go 1.16
GO := $(shell which go)
GO118 := $(shell which go1.18beta1)

build: prelude
	$(GO118) build

test: prelude
	$(GO118) test -v ./...

fmt: prelude
	$(GO118) fmt .

prelude:
	@$(GO118) version
	@date

go1.18:
	$(GO) install golang.org/dl/go1.18beta1@latest

ci: fmt test build

.PHONY: prelude test fmt build ci go1.18
