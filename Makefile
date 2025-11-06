.PHONY: build run tidy test clean lint

build:
	@go build -o bin/app cmd/note-taker/main.go

run: build
	@./bin/app

tidy: 
	@go mod tidy

test:
	@go test ./... -v

clean: 
	@rm -rf bin
lint:
	golangci-lint run ./...
