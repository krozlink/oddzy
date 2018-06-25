default: protoc test build

build:
	docker build --rm -t racing-service .

protoc:
	protoc --micro_out=. --go_out=. proto/racing.proto

test:
	go test -timeout 30s -v
