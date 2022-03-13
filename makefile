.PHONY: build test coverage

run:
	go run ./...

build:
	go build -o /tmp/fizz-buzz ./...

test:
	go test -v ./... --cover

lint:
	go vet ./...