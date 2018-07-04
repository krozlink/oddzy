
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker

proto:
	protoc --proto_path=${GOPATH}\src --proto_path=. --micro_out=. --go_out=. proto/racing.proto

build: proto

	go build -o racing-api main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t racing-api:latest
