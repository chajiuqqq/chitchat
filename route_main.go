package main

import (
	"github.com/chajiuqqq/chitchat/data"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func index(c *gin.Context) {
	threads, err := data.Threads()
	if err != nil {
		log.Error().Err(err).Msg("can't get threads")
		errorMsg(c, "can't get threads")
	} else {
		_,exist:=c.Get("sess")
		c.HTML(200, "index.tmpl", &gin.H{"IsPublic": !exist, "Threads": threads})
	}

}
