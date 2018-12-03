package client

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro"
	userpb "go-micro-demo1/user-service/proto/user"
	"context"
)

func GetTimeOrAddUser(name string,reg registry.Registry)(*userpb.UserTimeResponse,error)  {
	server := micro.NewService(
		micro.Name("go.micro.lzx.time"),
		micro.Registry(reg),
		micro.Version("latest"),
	)
	server.Init()

	uClient := userpb.NewUserService("go.micro.lzx.user", server.Client())
	tReq := &userpb.Username{
		Name:name,
	}
	return uClient.GetTimeOrAddUser(context.Background(), tReq)
}
