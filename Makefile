VERSION = 0.1.0
ROOT_DIR = $(CURDIR)

PLUGIN_PATH = $(CURDIR)/pkg/plugin

CONF_PATH = $(CURDIR)/pkg/internal/config
CONF_PROTO_FILES = "$(CONF_PATH)/*.proto"

API_PROTO_PATH = $(CURDIR)/pkg/api/copilot/v1
API_PROTO_FILES = "$(API_PROTO_PATH)/*.proto"

.PHONY: init
# init env
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	# go get -u github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2
	go get -u github.com/google/wire/cmd/wire
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install github.com/google/wire/cmd/wire
	npm install

.PHONY: api
# generate grpc code
api:
	cd $(API_PROTO_PATH) && protoc --proto_path=$(API_PROTO_PATH) \
           --proto_path=$(ROOT_DIR)/third_party \
           --go_out=paths=source_relative:. \
           --go-http_out=paths=source_relative:. \
		   --openapi_out=fq_schema_naming=true,default_response=false:. \
           $(API_PROTO_FILES)

.PHONY: proto
# generate internal proto struct
proto:
	cd $(CONF_PATH) && protoc --proto_path=$(CONF_PATH) \
           --proto_path=$(ROOT_DIR)/third_party \
           --go_out=paths=source_relative:. \
           $(CONF_PROTO_FILES)

.PHONY: wire
# generate wire
wire:
	cd $(PLUGIN_PATH) && wire

.PHONY: generate
# generate client code
generate:
	go generate ./...

.PHONY: build
# build
build:
	npm run build
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
all: api proto wire generate build run

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
