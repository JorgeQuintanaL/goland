build:
	@go build -C src/ -o ../bin/interfaces

run: build
	@./bin/interfaces

test:
	@go test -v ./...