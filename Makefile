NAME:=safekubectl

GO ?= go

.PHONY: all
all: build

#
# Go build related tasks
#
.PHONY: go-install
go-install:
	$(GO) install ./cmd/...

.PHONY: go-run
go-run: go-install
	$(GO) run ./cmd/* --v=4

.PHONY: go-fmt
go-fmt:
	gofmt -s -e -d $(shell find . -name "*.go" | grep -v /vendor/)

# Tests-related tasks
#
.PHONY: unit-test
unit-test: go-install
	go test -v ./pkg/...