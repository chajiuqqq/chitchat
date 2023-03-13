package rpc

import (
	"context"
	"fmt"
	"log"
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
		srvName := "authService"
		srv, err := NewDiscoveryClient().DiscoverService(srvName)
		if err != nil {
			log.Panic(err)
		}
		address := fmt.Sprintf("%s:%d", srv.Address, srv.Port)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Panic("connect error")
		}
		myAuthClient = pb.NewAuthServiceClient(conn)
	}()

	go func() {
		defer wg.Done()
		srvName := "threadService"
		srv, err := NewDiscoveryClient().DiscoverService(srvName)
		if err != nil {
			log.Panic(err)
		}
		address := fmt.Sprintf("%s:%d", srv.Address, srv.Port)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Panic("connect error")
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
