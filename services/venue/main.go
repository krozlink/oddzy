package main

import (
	"fmt"
	proto "github.com/krozlink/oddzy/services/venue/proto"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	fmt.Println("started")

	registry := consul.NewRegistry()

	srv := micro.NewService(
		micro.Name("venue"),
		micro.Version("latest"),
		micro.Registry(registry),
	)

	srv.Init()

	proto.RegisterVenueServiceHandler(srv.Server(), &service{})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
