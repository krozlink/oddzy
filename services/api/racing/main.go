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

	service.Init(
		micro.WrapHandler(RacingWrapper(service), CorrelationWrapper(service)),
	)

	proto.RegisterRacingHandler(service.Server(), new(Racing))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
