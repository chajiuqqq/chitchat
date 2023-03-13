package main

import (
	"fmt"
	"log"
	"net"

	"github.com/chajiuqqq/chitchat/common/pb"
	"github.com/chajiuqqq/chitchat/threadservice/service"
	"github.com/gin-gonic/gin"
	capi "github.com/hashicorp/consul/api"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

const (
	rpcPort  = 8000
	httpPort = 8001
	address  = "127.0.0.1"
)

func main() {
	group := new(errgroup.Group)

	group.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, rpcPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		threadService := new(service.ThreadService)
		pb.RegisterThreadServiceServer(grpcServer, threadService)
		return grpcServer.Serve(lis)
	})
	group.Go(func() error {
		// Get a new client
		client, err := capi.NewClient(capi.DefaultConfig())
		if err != nil {
			panic(err)
		}
		return registerService(client)
	})

	group.Go(func() error {
		r := gin.New()
		r.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, "ok")
		})
		return r.Run(fmt.Sprintf("%s:%d", address, httpPort))
	})

	// 等待所有 goroutine 完成
	if err := group.Wait(); err != nil {
		fmt.Println("Error:", err)
	}

}
func registerService(client *capi.Client) error {
	// 创建服务实例
	service := &capi.AgentServiceRegistration{
		Name:    "threadService",
		Port:    rpcPort,
		Address: address,
		Check: &capi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", address, httpPort),
			Interval: "10s",
			Timeout:  "2s",
		},
	}

	// 注册服务
	err := client.Agent().ServiceRegister(service)
	if err != nil {
		return err
	}

	return nil
}
