build:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	docker build --rm -t meeting-service .
	
localbuild:
	protoc -I. --proto_path=. --go_out=. --micro_out=. proto/*.proto
	go build -o ../../bin/meeting-service
	
run:
	docker run --rm -p 50051:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		meeting-service

localrun:
	../../bin/meeting-service
