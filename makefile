PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
NAME = $(shell basename $(PKG))
VERSION = $(shell cat .version)
COMMIT_SHA ?= $(shell git rev-parse --short HEAD)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CGO_ENABLED ?= 0

GOBUILD=CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -a -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

WORKSPACE ?= name

build:
	cd ./cmd/$(WORKSPACE) && $(GOBUILD)

# 生成swagger
swagger:
	swag init --pd -d ./cmd/$(WORKSPACE) -o ./cmd/$(WORKSPACE)/docs

# 生成openapi
openapi: swagger
	cd ./cmd/$(WORKSPACE) && klctl openapi  -f ./docs/swagger.json

