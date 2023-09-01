BASE_DIR=$(CURDIR)

.PHONY: tidy lean test

tidy:
	go mod tidy

format: tidy
	go fmt ./...

test: tidy
	go test -v ./...
