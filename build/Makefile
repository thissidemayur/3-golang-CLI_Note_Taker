.PHONY: build run tidy test clean lint

build:
	@go build -o bin/app .

run: build
	@./bin/app

tidy: 
	@go mod tidy

test:
	@go test ./... -v

clean: 
	@rm -rf bin
lint:
	@golangci-lint run ./...
