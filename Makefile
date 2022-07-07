.PHONY: build

PKG_LIST := $(shell go list ./... | grep -v /vendor/)
CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}

lint:
	golint -set_exit_status ${PKG_LIST}

swag-gen:
	echo ${REGISTRY}
	swag init -g api/routers.go -o api/docs

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --rebase

make create-env:
	cp ./.env.example ./.env