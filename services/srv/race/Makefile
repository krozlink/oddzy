build:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	docker build --rm -t race-service .
	
localbuild:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	go build -o ../../../bin/race-service
	
run:
	docker run --rm -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		race-service

localrun:
	../../../bin/race-service
