.PHONY: build cover start test test-integration

build:
	docker build -t truffle .

cover:
	go tool cover -html=cover.out

start:
	go run cmd/server/main.go

test:
	go test -coverprofile=cover.out -short ./...

test-integration:
	go test -coverprofile=cover.out -p 1 ./...


