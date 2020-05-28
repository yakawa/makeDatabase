GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

test:
	go generate ./...
	go test -v ./...

run:
	go generate ./...
	go run main.go

build:
	go generate ./...
	go build

cover:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html
	open cover.html
