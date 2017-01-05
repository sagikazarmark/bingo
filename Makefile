GO?=go
GOFMT?=gofmt
GLIDE?=glide
BINDATA?=go-bindata
GO_SOURCE_FILES = $(shell find . -type f -name "*.go" -not -name "bindata.go" -not -path "./vendor/*")

# install dependencies
install:
	@$(GLIDE) install

# generate resource files
generate:
	@$(GO) generate

# build the project and generate latest code
build:
	@$(GO) build

# check go sources
check:
	@$(GOFMT) -l ${GO_SOURCE_FILES} | read && echo "Code differs from gofmt's style" 1>&2 && exit 1 || true
	@$(GO) vet ${GO_SOURCE_FILES}

# gofmt excluding all the irrelevant go sources
fmt:
	@$(GOFMT) -l -w ${GO_SOURCE_FILES}

# fmt with simplification
fmts:
	@$(GOFMT) -l -s -w ${GO_SOURCE_FILES}

.PHONY: install generate build check fmt fmts
