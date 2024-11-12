build: 
	@go build -o bin/backend-api cmd/main.go

test: 
	@go test -v ./...

run: build 
	@./bin/backend-api