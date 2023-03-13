package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/chajiuqqq/chitchat/common/entity"
	"github.com/chajiuqqq/chitchat/common/pb"
	"github.com/chajiuqqq/chitchat/common/util"
	"github.com/gin-gonic/gin"
)

func newThread(c *gin.Context) {
	c.HTML(200, "newThread.tmpl", &gin.H{"IsPublic": false})
}

func createThread(c *gin.Context) {
	sess := c.MustGet("sess").(*entity.Session)
	rpcClient.MyThreadServiceClient.Create(context.Background(), &pb.CreateThreadReq{
		Topic:  c.PostForm("topic"),
		Uuid:   util.GenerateUuid(),
		UserId: uint32(sess.UserId),
	})
	c.Redirect(http.StatusFound, "/")

}

//对一个thread发表post
func postThread(c *gin.Context) {
	sess := c.MustGet("sess").(*entity.Session)
	tid, _ := strconv.Atoi(c.PostForm("id"))
	body := c.PostForm("body")
	_, err := rpcClient.MyThreadServiceClient.AddPost(context.Background(), &pb.AddPostRequest{
		ThreadId: uint32(tid),
		Body:     body,
		UserId:   uint32(sess.UserId),
	})
	if err != nil {
		log.Panic(err)
	}
	url := fmt.Sprintf("/thread/read/%d", tid)
	c.Redirect(http.StatusFound, url)
}

func readThread(c *gin.Context) {
	_, exist := c.Get("sess")
	tid, _ := strconv.Atoi(c.Param("tid"))
	getThreadRes, err := rpcClient.MyThreadServiceClient.Get(context.Background(), &pb.GetThreadRequest{
		ThreadId: uint32(tid),
	})
	if err != nil {
		log.Panic(err)
	}
	c.HTML(200, "thread.tmpl", &gin.H{"IsPublic": !exist, "Thread": getThreadRes})
}
