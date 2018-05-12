build:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	docker build -t venue-service .
	
localbuild:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	go build -o ../../bin/venue-service
	
run:
	docker run -p 50051:50051 -p 32771:8500 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		venue-service

localrun:
	../../bin/venue-service