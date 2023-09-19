BASE_DIR=$(CURDIR)

.PHONY: tidy lean test

tidy:
	go mod tidy

format: tidy
	go fmt ./...
	golines -w .

lint: format
	golangci-lint run

test: tidy
	go test -v ./...
