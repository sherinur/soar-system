BIN_NAME := "auth-bin"
MAIN_PATH := "cmd/auth/main.go"

build:
	go mod tidy
	gofumpt -l -w .
	go build -o $(BIN_NAME) $(MAIN_PATH)

run:
	gofumpt -l -w .																											
	go run $(MAIN_PATH)

clean:
	echo "Cleaning up..."
	rm -rf $(BIN_NAME)
	go mod tidy
	gofumpt -l -w .

upgrade:
	@echo "Cleaning the go mod cache..."
	go clean -modcache
	@echo "Upgrading the dependencies..."
	go get -u ./...

unit-test:
	@echo "Starting unit testing..."
	@go test ./internal/adapter/mongo/.
	@go test ./internal/adapter/grpc/server/frontend/.
	@go test ./internal/usecase/.