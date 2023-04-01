package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/chajiuqqq/chitchat/common/discover"
	"github.com/chajiuqqq/chitchat/common/pb"
	"github.com/chajiuqqq/chitchat/threadservice/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var consulService discover.DiscoveryClient = discover.NewConsulClient()
var (
	rpcPort  = flag.Int("rpcPort", 8000, "bind for RPC")
	httpPort = flag.Int("httpPort", 8001, "bind for http")
	address  = flag.String("address", "0.0.0.0", "bind ip for both RPC and http")
)

func main() {
	flag.Parse()
	group := new(errgroup.Group)

	group.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *address, *rpcPort))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Println("RPC listen on ", *address, ":", *rpcPort)
		grpcServer := grpc.NewServer()
		threadService := new(service.ThreadService)
		pb.RegisterThreadServiceServer(grpcServer, threadService)
		return grpcServer.Serve(lis)
	})
	group.Go(func() error {
		host := "threadService"
		err := consulService.Register("threadService", "", fmt.Sprintf("http://%s:%d/health", host, *httpPort), host, *rpcPort, nil, nil)
		return err
	})

	group.Go(func() error {
		r := gin.New()
		r.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, "ok")
		})
		return r.Run(fmt.Sprintf("%s:%d", *address, *httpPort))
	})

	// 等待所有 goroutine 完成
	if err := group.Wait(); err != nil {
		fmt.Println("Error:", err)
	}

}
