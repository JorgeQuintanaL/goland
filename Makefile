build:
	@go build -o bin/interfaces

run: build
	@./bin/interfaces

test:
	@go test -v ./...