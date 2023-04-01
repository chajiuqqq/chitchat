package main

import (
	"context"
	"log"

	"github.com/chajiuqqq/chitchat/frontend/utils"
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	res, err := threadClient.GetAll(context.Background())
	if err != nil {
		log.Println("can't get threads,", err)
		utils.ErrorMsg(c, "can't get threads"+err.Error())
		return
	}

	_, exist := c.Get("sess")
	c.HTML(200, "index.tmpl", &gin.H{"IsPublic": !exist, "Threads": res.Threads})

}
