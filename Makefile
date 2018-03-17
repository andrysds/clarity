build: pretty compile test

pretty:
	gofmt -w .

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build .

test:
	go test -v -race -cover -coverprofile coverage.out .
	go tool cover -html=coverage.out -o coverage.html
