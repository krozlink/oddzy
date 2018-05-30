build:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	go test -timeout 30s -v
	docker build --rm -t racing-service .

proto:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
