package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"MicroIhome/GetImageCd/handler"

	example "MicroIhome/GetImageCd/proto/example"
	"github.com/micro/go-grpc"
)

func main() {

	service := grpc.NewService(
		micro.Name("go.micro.srv.GetImageCd"),
		micro.Version("latest"),
	)

	service.Init()

	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
