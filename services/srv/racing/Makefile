default: proto test build

build:
	docker build --rm -t racing-service .

proto:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/racing.proto

test:
	go test -timeout 30s -v
