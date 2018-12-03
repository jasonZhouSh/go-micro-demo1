package client

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro"
	timePb "go-micro-demo1/time-service/proto/time"
	"context"
)

func GetUserTime(name string, reg registry.Registry)(*timePb.GetUserNameResponse, error)  {
	server := micro.NewService(
		micro.Name("go.micro.lzx.time"),
		micro.Registry(reg),
		micro.Version("latest"),
	)
	server.Init()

	tClient := timePb.NewTimeService("go.micro.lzx.time", server.Client())
	tReq := &timePb.Username{
		Name:name,
	}
	return tClient.GetUserTime(context.Background(), tReq)

}

func AddUser(name string,reg registry.Registry)(*timePb.AddUserResponse,error)  {

	server := micro.NewService(
		micro.Name("go.micro.lzx.time"),
		micro.Registry(reg),
		micro.Version("latest"),
	)
	server.Init()

	tClient := timePb.NewTimeService("go.micro.lzx.time", server.Client())
	tReq := &timePb.Username{
		Name:name,
	}
	return tClient.AddUser(context.Background(), tReq)
}