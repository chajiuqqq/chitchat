package client

import (
	"context"

	"github.com/chajiuqqq/chitchat/common/discover"
	"github.com/chajiuqqq/chitchat/common/loadbalance"
	"github.com/chajiuqqq/chitchat/common/pb"
)

type AuthClient interface {
	Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error)
	GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error)
	Encrypt(ctx context.Context, req *pb.EncryptRequest) (*pb.EncryptResponse, error)
	NewSession(ctx context.Context, req *pb.NewSessionReq) (*pb.AuthSession, error)
}

type AuthClientImpl struct {
	serviceName string
	manager     ClientManager
	loadBalance loadbalance.LoadBalance
}

func NewAuthClient(serviceName string, lb loadbalance.LoadBalance) AuthClient {
	if serviceName == "" {
		serviceName = "authservice"
	}
	if lb == nil {
		lb = DefaultLoadBalance
	}
	return &AuthClientImpl{
		serviceName: serviceName,
		manager: &DefaultManager{
			discoverClient: discover.ConsulService,
			loadBalance:    lb,
			serviceName:    serviceName,
		},
		loadBalance: lb,
	}
}
func (au *AuthClientImpl) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	res := &pb.CheckResponse{}
	err := au.manager.Invoke("/chitchat.AuthService/Check", "authservice", ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (au *AuthClientImpl) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	res := &pb.GetUserByEmailResponse{}
	err := au.manager.Invoke("/chitchat.AuthService/GetUserByEmail", "authservice", ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func(au *AuthClientImpl) Encrypt(ctx context.Context, req *pb.EncryptRequest) (*pb.EncryptResponse, error) {
	res := &pb.EncryptResponse{}
	err := au.manager.Invoke("/chitchat.AuthService/Encrypt", "authservice", ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, err
}

func (au *AuthClientImpl) NewSession(ctx context.Context, req *pb.NewSessionReq) (*pb.AuthSession, error) {
	res := &pb.AuthSession{}
	err := au.manager.Invoke("/chitchat.AuthService/NewSession", "authservice", ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, err
}
