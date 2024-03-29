// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: time.proto

/*
Package go_micro_lzx_time is a generated protocol buffer package.

It is generated from these files:
	time.proto

It has these top-level messages:
	User
	Username
	GetUserNameResponse
	AddUserResponse
*/
package go_micro_lzx_time

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for TimeService service

type TimeService interface {
	// 获取用户创建时间
	GetUserTime(ctx context.Context, in *Username, opts ...client.CallOption) (*GetUserNameResponse, error)
	// 添加新用户
	AddUser(ctx context.Context, in *Username, opts ...client.CallOption) (*AddUserResponse, error)
}

type timeService struct {
	c    client.Client
	name string
}

func NewTimeService(name string, c client.Client) TimeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.lzx.time"
	}
	return &timeService{
		c:    c,
		name: name,
	}
}

func (c *timeService) GetUserTime(ctx context.Context, in *Username, opts ...client.CallOption) (*GetUserNameResponse, error) {
	req := c.c.NewRequest(c.name, "TimeService.GetUserTime", in)
	out := new(GetUserNameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeService) AddUser(ctx context.Context, in *Username, opts ...client.CallOption) (*AddUserResponse, error) {
	req := c.c.NewRequest(c.name, "TimeService.AddUser", in)
	out := new(AddUserResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TimeService service

type TimeServiceHandler interface {
	// 获取用户创建时间
	GetUserTime(context.Context, *Username, *GetUserNameResponse) error
	// 添加新用户
	AddUser(context.Context, *Username, *AddUserResponse) error
}

func RegisterTimeServiceHandler(s server.Server, hdlr TimeServiceHandler, opts ...server.HandlerOption) {
	type timeService interface {
		GetUserTime(ctx context.Context, in *Username, out *GetUserNameResponse) error
		AddUser(ctx context.Context, in *Username, out *AddUserResponse) error
	}
	type TimeService struct {
		timeService
	}
	h := &timeServiceHandler{hdlr}
	s.Handle(s.NewHandler(&TimeService{h}, opts...))
}

type timeServiceHandler struct {
	TimeServiceHandler
}

func (h *timeServiceHandler) GetUserTime(ctx context.Context, in *Username, out *GetUserNameResponse) error {
	return h.TimeServiceHandler.GetUserTime(ctx, in, out)
}

func (h *timeServiceHandler) AddUser(ctx context.Context, in *Username, out *AddUserResponse) error {
	return h.TimeServiceHandler.AddUser(ctx, in, out)
}
