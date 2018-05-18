FROM golang:1.10.0 as builder

COPY . /go/src/github.com/krozlink/oddzy/services/venue
WORKDIR /go/src/github.com/krozlink/oddzy/services/venue

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o venue-service


FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/krozlink/oddzy/services/venue/venue-service .

CMD [ "./venue-service" ]