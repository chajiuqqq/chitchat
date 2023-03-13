package main

import (
	"context"
	"io"
	"log"

	"github.com/chajiuqqq/chitchat/common/pb"
	"github.com/chajiuqqq/chitchat/frontend/utils"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	stream, err := rpcClient.MyThreadServiceClient.GetAll(context.Background(), &pb.Empty{})
	if err != nil {
		log.Println("can't get threads,", err)
		utils.ErrorMsg(c, "can't get threads"+err.Error())
		return
	}
	threads := make([]*pb.GetThreadResponse, 0)
	for {
		item, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("fail to recv:", err)
		}
		threads = append(threads, item)
	}

	_, exist := c.Get("sess")
	c.HTML(200, "index.tmpl", &gin.H{"IsPublic": !exist, "Threads": threads})

}
