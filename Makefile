# ==================== #
# > > > Makefile < < < #
# ==================== #

# plugin version
VERSION = 0.1.0
# product root dir
ROOT_DIR = $(CURDIR)
# plugin tools dir
PLUGIN_TOOLS_DIR = $(CURDIR)/tool
PLUGIN_TOOL_PROTPC_GEN_GO_HTTP = $(PLUGIN_TOOLS_DIR)/protoc-gen-go-http
# plugin backend dir
PLUGIN_BACKEND_DIR = $(CURDIR)/pkg
PLUGIN_BACKEND_APP_DIR = $(PLUGIN_BACKEND_DIR)/plugin
# plugin backend config dir. config struct is designed at this dir
PLUGIN_BACKEND_CONFIG_DIR = $(PLUGIN_BACKEND_DIR)/internal/config
PLUGIN_BACKEND_CONFIG_FILES = "$(PLUGIN_BACKEND_CONFIG_DIR)/*.proto"
# plugin api dir
PLUGIN_BACKEND_API_DIR = $(PLUGIN_BACKEND_DIR)/api/copilot/v1
PLUGIN_BACKEND_API_FILES = "$(PLUGIN_BACKEND_API_DIR)/*.proto"

.PHONY: init
# init env
init:
	@echo "Init and install package dependencies and tools"
# install go package dependencies
	go mod tidy
# install protoc tool dependencies
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
# install protoc-gen-go grpc plugin: protoc-gen-go-grpc
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
# install protoc-gen-go http plugin: protoc-gen-go-http
	cd $(PLUGIN_TOOL_PROTPC_GEN_GO_HTTP) && go mod tidy && go install && cd $(ROOT_DIR)
# install wire [https://go-kratos.dev/docs/guide/wire]
	go get github.com/google/wire/cmd/wire@v0.6.0
	go install github.com/google/wire/cmd/wire
# install openapi tool
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
# install nodejs package dependencies
	npm install

.PHONY: api
# generate API code (http/grpc) 
api:
	@echo "Generate API"
	cd $(PLUGIN_BACKEND_API_DIR) && protoc --proto_path=$(PLUGIN_BACKEND_API_DIR) \
           --proto_path=$(ROOT_DIR)/third_party \
           --go_out=paths=source_relative:. \
           --go-http_out=paths=source_relative:. \
		   --openapi_out=fq_schema_naming=true,default_response=false:. \
           $(PLUGIN_BACKEND_API_FILES)

.PHONY: config
# generate internal config struct
config:
	@echo "Generate config struct"
	cd $(PLUGIN_BACKEND_CONFIG_DIR) && protoc --proto_path=$(PLUGIN_BACKEND_CONFIG_DIR) \
           --proto_path=$(ROOT_DIR)/third_party \
           --go_out=paths=source_relative:. \
           $(PLUGIN_BACKEND_CONFIG_FILES)

.PHONY: wire
# generate wire_gen.go
wire:
	@echo "Generate wire_gen.go"
	cd $(PLUGIN_BACKEND_APP_DIR) && wire

.PHONY: generate
# generate client code
generate:
	@echo "Generate client code"
	go generate ./...

.PHONY: build-frontend
# build frontend
build-frontend:
	@echo "Build frontend"
	npm run build

.PHONY: build-backend
# build backend
build-backend:
	@echo "Build backend"
	mage -v

.PHONY: build
# build all
build: 
	@echo "Build All"
	build-frontend 
	build-backend

.PHONY: test
# backend test
test:
	@echo "Backend test"
	go test -v ./... -cover

.PHONY: run
# docker compose run
run:
	@echo "Docker compose run"
	docker-compose.exe up


.PHONY: all
# workflow here
all: 
	@echo "Workflow executing"
	init 
	api 
	config 
	wire 
	generate 
	build 
	run
	@echo "Workflow finished without error"

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
