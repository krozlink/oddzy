FROM golang:1.10.0 as builder

COPY . /go/src/github.com/krozlink/oddzy/services/meeting
WORKDIR /go/src/github.com/krozlink/oddzy/services/meeting

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o meeting-service


FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/krozlink/oddzy/services/meeting/meeting-service .

CMD [ "./meeting-service" ]