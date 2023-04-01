package client

import (
	"context"

	"github.com/chajiuqqq/chitchat/common/discover"
	"github.com/chajiuqqq/chitchat/common/loadbalance"
	"github.com/chajiuqqq/chitchat/common/pb"
)

type ThreadClient interface {
	Get(ctx context.Context, req *pb.GetThreadRequest) (*pb.GetThreadResponse, error)
	GetAll(ctx context.Context) (*pb.GetAllResponse, error)
	Create(ctx context.Context, req *pb.CreateThreadReq) error
	AddPost(ctx context.Context, req *pb.AddPostRequest) error
}

type ThreadClientImpl struct {
	serviceName string
	manager     ClientManager
	loadBalance loadbalance.LoadBalance
}

func NewThreadClient(serviceName string, lb loadbalance.LoadBalance) ThreadClient {
	if serviceName == "" {
		serviceName = "threadservice"
	}
	if lb == nil {
		lb = DefaultLoadBalance
	}
	return &ThreadClientImpl{
		serviceName: serviceName,
		manager: &DefaultManager{
			discoverClient: discover.ConsulService,
			loadBalance:    lb,
			serviceName:    serviceName,
		},
		loadBalance: lb,
	}
}

func (tc *ThreadClientImpl) Get(ctx context.Context, req *pb.GetThreadRequest) (*pb.GetThreadResponse, error) {
	res := &pb.GetThreadResponse{}
	err := tc.manager.Invoke("/chitchat.ThreadService/Get", "threadservice", ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (tc *ThreadClientImpl) GetAll(ctx context.Context) (*pb.GetAllResponse, error) {
	res := &pb.GetAllResponse{}
	err := tc.manager.Invoke("/chitchat.ThreadService/GetAll", "threadservice", ctx, nil, res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (tc *ThreadClientImpl) Create(ctx context.Context, req *pb.CreateThreadReq) error {
	err := tc.manager.Invoke("/chitchat.ThreadService/Create", "threadservice", ctx, req, &pb.Empty{})
	if err != nil {
		return err
	}
	return nil
}

func (tc *ThreadClientImpl) AddPost(ctx context.Context, req *pb.AddPostRequest) error {
	err := tc.manager.Invoke("/chitchat.ThreadService/AddPost", "threadservice", ctx, req, &pb.Empty{})
	if err != nil {
		return err
	}
	return nil
}
