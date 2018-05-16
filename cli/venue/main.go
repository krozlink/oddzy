package main

import (
	"context"
	"fmt"
	venue "github.com/krozlink/oddzy/services/venue/proto"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService()

	service.Init()

	cli := venue.NewVenueService("oddzy.services.venue", service.Client())

	rsp, err := cli.List(context.Background(), &venue.ListRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Venues)
}
