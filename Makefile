.PHONY: usage build get-linter lint

OK_COLOR=\033[32;01m
NO_COLOR=\033[0m
WARN_COLOR=\033[33;01m
ERROR_COLOR=\033[31;01m

GO := go
GO_LINTER := golint
ECHOFLAGS ?=
GOFLAGS ?=

BUILDOS ?= $(shell go env GOHOSTOS)
BUILDARCH ?= amd64

ENVFLAGS ?= CGO_ENABLED=0
BUILDENV ?= GOOS=$(BUILDOS) GOARCH=$(BUILDARCH)

BIN_COURSE := course

PKGS = $(shell $(GO) list ./...)

usage: Makefile
	@echo $(ECHOFLAGS) "to use make call:"
	@echo $(ECHOFLAGS) "    make <action>"
	@echo $(ECHOFLAGS) ""
	@echo $(ECHOFLAGS) "list of available actions:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## build: build course binary.
build: lint staticcheck
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Building binary (linux/amd64/$(BIN_COURSE))...$(NO_COLOR)"
	@echo $(ECHOFLAGS) $(ENVFLAGS) $(BUILDENV) $(GO) build $(BUILDFLAGS) -o bin/linux_amd64/$(BIN_COURSE) ./cmd
	@$(ENVFLAGS) $(BUILDENV) $(GO) build -o bin/linux_amd64/$(BIN_COURSE) ./cmd

## staticcheck: run staticcheck on packages
staticcheck:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running staticcheck...$(NO_COLOR)"
	@$(GO) get -v honnef.co/go/tools/cmd/staticcheck
	@$(ENVFLAGS) staticcheck $(PKGS)

## get-linter: install linter
get-linter:
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Getting linter...$(NO_COLOR)"
	@go get -v -u golang.org/x/lint/golint

## lint: lint package
lint: get-linter
	@echo $(ECHOFLAGS) "$(OK_COLOR)==> Running linter...$(NO_COLOR)"
	@$(GO_LINTER) -set_exit_status $(PKGS)