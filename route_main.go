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
		_, err = SessionCheck(c.Writer, c.Request)
		c.HTML(200, "index.tmpl", &gin.H{"IsPublic": err != nil, "Threads": threads})
	}

}
