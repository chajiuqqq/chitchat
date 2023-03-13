package rpc

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/chajiuqqq/chitchat/common/entity"
	"github.com/chajiuqqq/chitchat/common/pb"
	"github.com/gin-gonic/gin"
	capi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
)

type Consumer interface {
	Reload(srv *capi.AgentService) (done chan struct{}, err error)
}

type rpcClient struct {
	MyAuthServiceClient   pb.AuthServiceClient
	MyThreadServiceClient pb.ThreadServiceClient
}

var myRpcCient *rpcClient

func NewRpcClient() *rpcClient {
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
	NewDiscoveryClient().HealthCheck(myRpcCient)
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


//服务地址发生变化时，reload当前rpc连接
func (rpc *rpcClient) Reload(srv *capi.AgentService) (done chan struct{}, err error) {
	done = make(chan struct{})
	log.Println("reload ", srv.Service, "new port:", srv.Port)
	switch srv.Service {
	case "authService":
		defer close(done)
		address := fmt.Sprintf("%s:%d", srv.Address, srv.Port)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Panic(err)
		}
		rpc.MyAuthServiceClient = pb.NewAuthServiceClient(conn)

	case "threadService":

		defer close(done)
		address := fmt.Sprintf("%s:%d", srv.Address, srv.Port)
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Panic(err)
		}
		rpc.MyThreadServiceClient = pb.NewThreadServiceClient(conn)

	}
	return
}
