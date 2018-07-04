package main

import (
	proto "github.com/krozlink/oddzy/services/api/racing/proto"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.racing"),
		micro.Version("latest"),
	)

	service.Init()

	proto.RegisterRacingHandler(service.Server(), new(handler))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
