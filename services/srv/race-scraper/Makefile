build:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	docker build --rm -t race-scraper-service .
	
run:
	docker run --rm -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		venrace-race-scraper-service


test:
	go test -timeout 30s -v
