default: protoc test build

build:
	docker build --rm -t race-scraper-service .

protoc:
	protoc --micro_out=. --go_out=. proto/scraper.proto


run:
	docker run --rm -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		venrace-race-scraper-service


test:
	go test -timeout 30s -v
