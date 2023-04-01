package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/chajiuqqq/chitchat/common/client"
	_ "github.com/chajiuqqq/chitchat/common/data"
	"github.com/chajiuqqq/chitchat/common/discover"
	"github.com/chajiuqqq/chitchat/common/entity"
	"github.com/chajiuqqq/chitchat/common/pb"
	_ "github.com/chajiuqqq/chitchat/common/pb"
	"github.com/chajiuqqq/chitchat/frontend/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var authClient = client.NewAuthClient("authservice", client.DefaultLoadBalance)
var threadClient = client.NewThreadClient("threadservice", client.DefaultLoadBalance)
var consulService discover.DiscoveryClient = discover.NewConsulClient()

func sessionCheck(c *gin.Context) (sess *entity.Session, err error) {
	cookie, err := c.Cookie("_cookie")
	if err != nil {
		return
	}
	checkResponse, err := authClient.Check(context.Background(), &pb.CheckRequest{
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

func myLoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess, err := sessionCheck(c)
		if err == nil {
			c.Set("sess", sess)
		}
		c.Next()
	}
}

func myAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exist := c.Get("sess"); !exist {
			c.Redirect(http.StatusFound, "/frontend/login")
			return
		}
		c.Next()
	}
}
func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.ErrorMsg(c, err.(error).Error())
			}
			switch c.Writer.Status() {
			case 200:
				return
			case 404:
				utils.ErrorMsg(c, "404 Not Found: "+c.Request.URL.Path)

			}
		}()
		c.Next()
	}
}

var (
	port = flag.Int("httpPort", 8080, "bind http")
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func main() {
	flag.Parse()
	log.Println("Chitchat start at %d", *port)
	r := gin.New()
	r.Use(gin.Logger(), myLoginCheck(), recovery())
	r.SetFuncMap(template.FuncMap{
		"timeFormat": func(t time.Time) string {
			return t.Format("2006.01.02 15:04:05")
		},
		"NumReplies": func(th *pb.GetThreadResponse) int {
			if th.Posts != nil {
				return len(th.Posts)
			}
			return 0
		},
	})
	r.Static("/static", "../public")
	r.LoadHTMLGlob("../templates/**/*")
	r.GET("/", index)
	r.GET("/err", err)
	r.GET("/login", login)
	r.GET("/logout", logout)
	r.GET("/signup", signup)

	r.POST("/signup_account", signupAccount)
	r.POST("/authenticate", authenticate)
	r.GET("/thread/read/:tid", readThread)
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, "ok")
	})

	threadGroup := r.Group("/thread")
	threadGroup.Use(myAuth())
	threadGroup.GET("/new", newThread)
	threadGroup.POST("/create", createThread)
	threadGroup.POST("/post", postThread)

	group := new(errgroup.Group)

	group.Go(func() error {
		host := "frontend"
		err := consulService.Register("frontend", "", fmt.Sprintf("http://%s:%d/health", host, *port), host, *port, nil, nil)
		return err
	})

	group.Go(func() error {

		return r.Run(fmt.Sprintf(":%d", *port))
	})
	// 等待所有 goroutine 完成
	if err := group.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
}
