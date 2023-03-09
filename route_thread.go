package main

import (
	"fmt"
	"net/http"

	"github.com/chajiuqqq/chitchat/data"
	"github.com/gin-gonic/gin"
)

func newThread(c *gin.Context) {
	c.HTML(200, "newThread.tmpl", &gin.H{"IsPublic": false})
}

func createThread(c *gin.Context) {
	sess := c.MustGet("sess").(*data.Session)
	thread := data.Thread{
		Topic:  c.PostForm("topic"),
		Uuid:   generateUuid(),
		UserId: sess.UserId,
	}
	data.Db.Create(&thread)
	c.Redirect(http.StatusFound, "/")

}

//对一个thread发表post
func postThread(c *gin.Context) {
	sess := c.MustGet("sess").(*data.Session)
	uuid := c.PostForm("uuid")
	body := c.PostForm("body")
	var thread data.Thread
	data.Db.Where("uuid=?", uuid).First(&thread)
	data.Db.Model(&thread).Association("Posts").Append(
		&data.Post{Uuid: generateUuid(), Body: body, UserId: sess.UserId},
	)

	url := fmt.Sprintf("/thread/read/%s", uuid)
	c.Redirect(http.StatusFound, url)
}

func readThread(c *gin.Context) {
	_, exist := c.Get("sess")
	tid := c.Param("tid")
	thread, err := data.GetThread(tid)
	if err != nil {
		errorMsg(c, err.Error())
	}
	c.HTML(200, "thread.tmpl", &gin.H{"IsPublic": !exist, "Thread": thread})
}
