package server

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/pkg/errors"
	timepb "go-micro-demo1/time-service/proto/time"
	"log"
	"time"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/go-micro/registry"
)

// 定义服务
type service struct {
	repo Repository
}

// 实现服务端
func (s *service) GetUserTime(ctx context.Context, uname *timepb.Username, resp *timepb.GetUserNameResponse) error {
	// 调用内部方法
	user, err := s.repo.GetUserTime(uname)
	if err != nil {
		return err
	}
	resp.User = user
	return nil
}
func (s *service) AddUser(ctx context.Context, uname *timepb.Username, resp *timepb.AddUserResponse) error {
	// 调用内部方法
	user, err := s.repo.AddUser(uname)
	if err != nil {
		return err
	}
	resp.User = user
	return nil
}

type Repository interface {
	GetUserTime(*timepb.Username) (*timepb.User, error)
	AddUser(*timepb.Username) (*timepb.User, error)
}

type timeRepository struct {
	userMap map[string]*timepb.User
}

// 接口实现
func (repo *timeRepository) GetUserTime(uname *timepb.Username) (*timepb.User, error) {
	user, ok := repo.userMap[uname.Name]
	if ok {
		return user, nil
	}
	return nil, errors.New("GetUserTime --- the user is not exist.")
}

// 接口实现
func (repo *timeRepository) AddUser(uname *timepb.Username) (*timepb.User, error) {

	_, ok := repo.userMap[uname.Name]
	if ok {
		return nil, errors.New("AddUser --- the user is exist.")
	}
	nowtime := time.Now().Unix()
	repo.userMap[uname.Name] = &timepb.User{
		Id:        fmt.Sprintf("id%ld", nowtime),
		Name:      uname.Name,
		CrateTime: int32(nowtime),
	}
	return &timepb.User{
		Id:        fmt.Sprintf("id%ld", nowtime),
		Name:      uname.Name,
		CrateTime: int32(nowtime),
	}, nil
}

func timeserver() {
	//连接etcd
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"http://127.0.0.1:2379",
		}
	})


	server := micro.NewService(
		micro.Name("go.micro.lzx.time"),
		micro.Registry(reg),
		micro.Version("latest"),
	)
	server.Init()

	repo := &timeRepository{
		userMap: make(map[string]*timepb.User),
	}

	//3 将实现服务端的 API 注册到服务端
	timepb.RegisterTimeServiceHandler(server.Server(), &service{repo})

	//4 服务端运行
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
