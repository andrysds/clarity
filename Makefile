build: pretty compile test

pretty:
	gofmt -w .

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build .

test:
	./test.sh
	go tool cover -html=coverage.out -o coverage.html

dep:
	dep ensure -vendor-only

doc:
	godoc -http=:6060 -q
