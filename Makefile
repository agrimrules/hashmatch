PACKAGE := go.agrim.dev/hashmatch
VERSION ?= $(shell git describe --abbrev=0 --tags)

default: build

build:
	@go build \
	-ldflags "-w -s -X ${PACKAGE}/cmd.version=${VERSION}" \
	-a .

test:
	@echo "coming soon.."