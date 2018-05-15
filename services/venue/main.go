package main

import (
	"fmt"
	proto "github.com/krozlink/oddzy/services/venue/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
)

func main() {
	fmt.Println("started")
	srv := micro.NewService(
		micro.Name("oddzy.services.venue"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterVenueServiceHandler(srv.Server(), &service{})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
