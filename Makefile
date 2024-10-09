VERSION = 0.1.0
ROOT_DIR = $(CURDIR)
VVAULT_PLUGIN_PATH = $(CURDIR)/pkg/plugin
VVAULT_CONF_PATH = $(CURDIR)/pkg/internal/config
VVAULT_CONF_PROTO_FILES = "$(VVAULT_CONF_PATH)/*.proto"

.PHONY: init
# init env
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/google/wire/cmd/wire
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/google/wire/cmd/wire

.PHONY: proto
# generate internal proto struct
proto:
	cd $(VVAULT_CONF_PATH) && protoc --proto_path=$(VVAULT_CONF_PATH) \
           --proto_path=$(ROOT_DIR)/third_party \
           --go_out=paths=source_relative:. \
           $(VVAULT_CONF_PROTO_FILES)

.PHONY: wire
# generate wire
wire:
	cd $(VVAULT_PLUGIN_PATH) && wire

.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: build
# build
build:
	npm install
	mage -v

.PHONY: test
# test
test:
	go test -v ./... -cover

.PHONY: run
run:
	docker-compose.exe up


.PHONY: all
# generate all
all: proto generate build test run

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
