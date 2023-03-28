package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	_ "github.com/chajiuqqq/chitchat/common/data"
	"github.com/chajiuqqq/chitchat/common/pb"
	_ "github.com/chajiuqqq/chitchat/common/pb"
	"github.com/chajiuqqq/chitchat/common/rpc"
	"github.com/chajiuqqq/chitchat/frontend/utils"
	"github.com/gin-gonic/gin"
	capi "github.com/hashicorp/consul/api"
	"golang.org/x/sync/errgroup"
)

var rpcClient = rpc.NewRpcClient()

func myLoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess, err := rpcClient.SessionCheck(c)
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
		// Get a new client
		config := capi.DefaultConfig()
		config.Address = "consul:8500"
		client, err := capi.NewClient(config)
		if err != nil {
			panic(err)
		}
		return registerService(client)
	})

	group.Go(func() error {

		return r.Run(fmt.Sprintf(":%d", *port))
	})
	// 等待所有 goroutine 完成
	if err := group.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
}

func registerService(client *capi.Client) error {
	host := "frontend"
	// 创建服务实例
	service := &capi.AgentServiceRegistration{
		Name:    "frontend",
		Port:    *port,
		Address: host,
		Check: &capi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%d/health", host, *port),
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
