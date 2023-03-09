package main

import (
	"net/http"
	"text/template"
	"time"

	_ "github.com/chajiuqqq/chitchat/data"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func myLoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess, err := SessionCheck(c)
		if err == nil {
			c.Set("sess", sess)
		}
		c.Next()
	}
}

func myAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exist := c.Get("sess"); !exist {
			c.Redirect(http.StatusFound, "/login")
		}
		c.Next()
	}
}
func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errorMsg(c, err.(error).Error())
			}
			switch c.Writer.Status() {
			case 200:
				return
			case 404:
				errorMsg(c, "404 Not Found")

			}
		}()
		c.Next()
	}
}

func main() {
	log.Info().Msgf("Chitchat %s start at %s", Config.Version, Config.Port)
	r := gin.New()
	r.Use(gin.Logger(), myLoginCheck(), recovery())
	r.SetFuncMap(template.FuncMap{
		"timeFormat": func(t time.Time) string {
			return t.Format("2006.01.02 15:04:05")
		},
	})
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./public")
	r.GET("/", index)
	r.GET("/err", err)
	r.GET("/login", login)
	r.GET("/logout", logout)
	r.GET("/signup", signup)

	r.POST("/signup_account", signupAccount)
	r.POST("/authenticate", authenticate)
	r.GET("/thread/read/:tid", readThread)

	threadGroup := r.Group("/thread")
	threadGroup.Use(myAuth())
	threadGroup.GET("/new", newThread)
	threadGroup.POST("/create", createThread)
	threadGroup.POST("/post", postThread)
	r.Run(":8080")
}
