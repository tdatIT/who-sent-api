#setup > wire > clean > build > run

SERVICE_NAME = example
WORKER_MAIN_FILE = server_app
BUILD_DIR = $(PWD)/build
GO= go

setup:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest

wire:
	cd internal/ && wire

protoc-infra:
	protoc \
      -I internal/infrastructure/rpcClient/ \
      -I internal/infrastructure/rpcClient/third-party/validate \
      --go_out=internal/infrastructure/rpcClient/pb/ \
      --go_opt=paths=source_relative \
      --go-grpc_out=internal/infrastructure/rpcClient/pb/ \
      --go-grpc_opt=paths=source_relative \
      --validate_out="lang=go,paths=source_relative:internal/infrastructure/rpcClient/pb" \
      internal/infrastructure/rpcClient/*.proto

test:
	go test ./...

# clean build file
clean:
	echo "remove bin exe"
	rm -rf $(BUILD_DIR)

lint:
	golangci-lint run ./...
# build binary
build-bin:
	echo "build binary execute file"
	make wire
	cd cmd/ && GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(WORKER_MAIN_FILE)_linux .

run:
	make build
	echo "Run service application"
	cd $(BUILD_DIR) && ./$(WORKER_MAIN_FILE)_linux

