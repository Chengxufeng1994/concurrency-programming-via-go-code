# Go compiler selection
ifeq ($(GO),)
GO := go
endif

GOROOT := $(shell "$(GO)" env GOROOT)
export GOROOT := $(GOROOT)
GO_EXE_PATH := $(GOROOT)/bin

# Make sure we are in the same directory as the Makefile
BASE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

TOOLS_DIR=tools

# PATH
export PATH := $(BASE_DIR)/$(TOOLS_DIR):$(GO_EXE_PATH):$(PATH)

# Kernel (OS) Name
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')

# Allow architecture to be overwritten
ifeq ($(HOST_ARCH),)
HOST_ARCH := $(shell uname -m)
endif

# Build architecture settings:
# EXEC_ARCH defines the architecture of the executables that gets compiled
# DOCKER_ARCH defines the architecture of the docker image
# Both vars must be set, an unknown architecture defaults to amd64
ifeq (x86_64, $(HOST_ARCH))
EXEC_ARCH := amd64
DOCKER_ARCH := amd64
else ifeq (i386, $(HOST_ARCH))
EXEC_ARCH := 386
DOCKER_ARCH := i386
else ifneq (,$(filter $(HOST_ARCH), arm64 aarch64))
EXEC_ARCH := arm64
DOCKER_ARCH := arm64
else ifeq (armv7l, $(HOST_ARCH))
EXEC_ARCH := arm
DOCKER_ARCH := arm32v7
else
$(info Unknown architecture "${HOST_ARCH}" defaulting to: amd64)
EXEC_ARCH := amd64
DOCKER_ARCH := amd64
endif

# golangci-lint
GOLANGCI_LINT_VERSION=1.62.2
GOLANGCI_LINT_PATH=$(TOOLS_DIR)/golangci-lint-v$(GOLANGCI_LINT_VERSION)
GOLANGCI_LINT_BIN=$(GOLANGCI_LINT_PATH)/golangci-lint
GOLANGCI_LINT_ARCHIVE=golangci-lint-$(GOLANGCI_LINT_VERSION)-$(OS)-$(EXEC_ARCH).tar.gz
GOLANGCI_LINT_ARCHIVEBASE=golangci-lint-$(GOLANGCI_LINT_VERSION)-$(OS)-$(EXEC_ARCH)
export PATH := $(BASE_DIR)/$(GOLANGCI_LINT_PATH):$(PATH)

# Install golangci-lint
$(GOLANGCI_LINT_BIN):
	@echo "installing golangci-lint v$(GOLANGCI_LINT_VERSION)"
	@mkdir -p "$(GOLANGCI_LINT_PATH)"
	@curl -sSfL "https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/$(GOLANGCI_LINT_ARCHIVE)" \
		| tar -x -z --strip-components=1 -C "$(GOLANGCI_LINT_PATH)" "$(GOLANGCI_LINT_ARCHIVEBASE)/golangci-lint"

# Run lint against the previous commit for PR and branch build
# In dev setup look at all changes on top of master
.PHONY: lint
lint: $(GOLANGCI_LINT_BIN)
	@echo "running golangci-lint"
	@"${GOLANGCI_LINT_BIN}" run