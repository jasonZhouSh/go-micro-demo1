package server

import (
	"context"
	"github.com/micro/go-micro"
	userpb "go-micro-demo1/user-service/proto/user"
	"log"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
	"go-micro-demo1/user-service/client"
	"fmt"
	"github.com/golang/glog"
)

// 定义服务
type service struct {
	repo Repository
}

// 实现服务端
func (s *service) GetTimeOrAddUser(ctx context.Context, uname *userpb.Username, resp *userpb.UserTimeResponse) error {
	// 调用内部方法
	time, err := s.repo.GetTimeOrAddUser(uname)
	if err != nil {
		return err
	}
	resp.Time = time
	return nil
}


type Repository interface {
	GetTimeOrAddUser(*userpb.Username) (int32, error)
}

type userRepository struct {
}

// 接口实现
func (repo *userRepository) GetTimeOrAddUser(uname *userpb.Username) (int32, error) {
	glog.Info("GetTimeOrAddUser...")
	glog.Info("uname.Name:",uname.Name)
	//向time服务请求创建,有则直接返回时间
	urep,err:=client.GetUserTime(uname.Name,reg)
	if err==nil  {
		glog.Info("urep.User:",urep.User.Name)
		glog.Info("urep.User:",urep.User.CrateTime)
		return urep.User.CrateTime,nil
	}

	user,err:=client.AddUser(uname.Name,reg)
	if err!=nil {
		fmt.Println(err)
		return 0, err
	}

	glog.Info("user.User.CrateTime:",user.User.CrateTime)
	return user.User.CrateTime,nil
}

var reg registry.Registry
func userserver() {
	//连接etcd
	reg = etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})

	server := micro.NewService(
		micro.Name("go.micro.lzx.user"),
		micro.Registry(reg),
		micro.Version("latest"),
	)
	server.Init()
	repo := &userRepository{
	}

	//3 将实现服务端的 API 注册到服务端
	userpb.RegisterUserServiceHandler(server.Server(), &service{repo})

	//4 服务端运行
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
