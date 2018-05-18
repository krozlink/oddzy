FROM golang:1.10.0 as builder

COPY . /go/src/github.com/krozlink/oddzy/services/race
WORKDIR /go/src/github.com/krozlink/oddzy/services/race

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o race-service


FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/krozlink/oddzy/services/race/race-service .

CMD [ "./race-service" ]