.PHONY: build
build:
	go build -v ./cmd/appServer

.PHONY: test
test:
	go test -v -race ./...

.PHONY: buildWindows
buildWindows:
	GOOS=windows GOARCH=amd64 go build -v ./cmd/appServer

.DEFAULT_GOAL := build