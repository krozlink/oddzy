FROM golang:1.10.0 as builder

COPY . /go/src/github.com/krozlink/oddzy/services/api/racing
WORKDIR /go/src/github.com/krozlink/oddzy/services/api/racing

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o racing-api


FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/krozlink/oddzy/services/api/racing/racing-api .

CMD [ "./racing-api" ]