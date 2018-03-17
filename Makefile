all : pretty test build

pretty:
	gofmt -w .

test:
	go test -v -race -cover -coverprofile coverage.out .
	go tool cover -html=coverage.out -o coverage.html

build:
	go build .
