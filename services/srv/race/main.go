package main

import (
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/race/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
)

func main() {
	fmt.Println("started")
	srv := micro.NewService(
		micro.Name("oddzy.services.race"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterRaceServiceHandler(srv.Server(), &service{})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
