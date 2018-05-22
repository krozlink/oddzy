build:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	docker build --rm -t race-scraper-service .
	
localbuild:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	go build -o ../../../bin/race-scraper-service
	
run:
	docker run --rm -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		venrace-race-scraper-service

localrun:
	../../../bin/race-scraper-service
