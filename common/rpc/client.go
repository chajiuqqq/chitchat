package rpc

import (
	"context"
	"sync"

	"github.com/chajiuqqq/chitchat/common/entity"
	"github.com/chajiuqqq/chitchat/common/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type rpcClient struct {
	MyAuthServiceClient   pb.AuthServiceClient
	MyThreadServiceClient pb.ThreadServiceClient
}

var myRpcCient *rpcClient

func New() *rpcClient {
	if myRpcCient != nil {
		return myRpcCient
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var myAuthClient pb.AuthServiceClient
	var myThreadServiceClient pb.ThreadServiceClient

	go func() {
		defer wg.Done()
		serviceAddress := "127.0.0.1:8100"
		conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
		if err != nil {
			panic("connect error")
		}
		myAuthClient = pb.NewAuthServiceClient(conn)
	}()

	go func() {
		defer wg.Done()
		serviceAddress := "127.0.0.1:8000"
		conn, err := grpc.Dial(serviceAddress, grpc.WithInsecure())
		if err != nil {
			panic("connect error")
		}
		myThreadServiceClient = pb.NewThreadServiceClient(conn)
	}()

	wg.Wait()
	myRpcCient = &rpcClient{
		MyAuthServiceClient:   myAuthClient,
		MyThreadServiceClient: myThreadServiceClient,
	}
	return myRpcCient

}

func (rpc *rpcClient) SessionCheck(c *gin.Context) (sess *entity.Session, err error) {
	cookie, err := c.Cookie("_cookie")
	if err != nil {
		return
	}
	checkResponse, err := rpc.MyAuthServiceClient.Check(context.Background(), &pb.CheckRequest{
		Uuid: cookie,
	})
	if err != nil || !checkResponse.Exist {
		return
	}
	return &entity.Session{
		Uuid:   cookie,
		Email:  checkResponse.Sess.Email,
		UserId: uint(checkResponse.Sess.UserId),
	}, nil
}
