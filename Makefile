# Build the Go application and generate the binary in the bin directory
build:
	@go build -o bin/backend-api cmd/main.go

# Run the tests in all subdirectories
test:
	@go test -v ./...

# Run the compiled application after building
run: build
	@./bin/backend-api