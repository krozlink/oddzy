FROM golang:1.10.0 as builder

COPY . /go/src/github.com/krozlink/oddzy/services/srv/race-scraper
WORKDIR /go/src/github.com/krozlink/oddzy/services/srv/race-scraper

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o race-scraper-service


FROM alpine:latest

RUN mkdir /app
WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/src/github.com/krozlink/oddzy/services/srv/race-scraper/race-scraper-service .

CMD [ "./race-scraper-service" ]