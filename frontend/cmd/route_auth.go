package main

import (
	"context"
	"log"
	"net/http"

	"github.com/chajiuqqq/chitchat/common/data"
	"github.com/chajiuqqq/chitchat/common/entity"
	"github.com/chajiuqqq/chitchat/common/pb"
	"github.com/chajiuqqq/chitchat/common/util"
	"github.com/chajiuqqq/chitchat/frontend/utils"
	"github.com/gin-gonic/gin"
)

func authenticate(c *gin.Context) {
	getUserByEmailResponse, err := rpcClient.MyAuthServiceClient.GetUserByEmail(context.Background(), &pb.GetUserByEmailRequest{
		Email: c.PostForm("email"),
	})
	if err != nil {
		panic(err)
	}

	encryptResponse, err := rpcClient.MyAuthServiceClient.Encrypt(context.Background(), &pb.EncryptRequest{
		Src: c.PostForm("password"),
	})

	if err != nil {
		panic(err)
	}

	if !getUserByEmailResponse.Exist {
		c.Redirect(http.StatusFound, "/login")
	} else if user := getUserByEmailResponse.User; user.Password == encryptResponse.Out {
		session, err := rpcClient.MyAuthServiceClient.NewSession(context.Background(), &pb.NewSessionReq{
			UserId: user.ID,
			Email:  user.Email,
		})
		if err != nil {
			log.Panic("can't get new session,", err)
		}
		c.SetCookie("_cookie", session.Uuid, 3153600, "/", "", true, true)
		c.Set("sess", session)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.Redirect(http.StatusFound, "/login")
	}

}

func login(c *gin.Context) {
	c.HTML(200, "login.tmpl", gin.H{"IsPublic": true})
}

func logout(c *gin.Context) {
	c.SetCookie("_cookie", "", -1, "/", "", true, true)
	c.Redirect(http.StatusFound, "/")
}

func signup(c *gin.Context) {
	c.HTML(200, "signup.tmpl", gin.H{"IsPublic": true})
}

func signupAccount(c *gin.Context) {
	encryptResponse, err := rpcClient.MyAuthServiceClient.Encrypt(context.Background(), &pb.EncryptRequest{
		Src: c.PostForm("password"),
	})
	if err != nil {
		utils.ErrorMsg(c, err.Error())
		return
	}
	user := entity.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: encryptResponse.Out,
		Uuid:     util.GenerateUuid(),
	}
	data.Db.Create(&user)
	c.Redirect(http.StatusFound, "/")
}

func err(c *gin.Context) {
	_, exist := c.Get("sess")
	c.HTML(200, "error.tmpl", gin.H{"IsPublic": !exist, "Msg": c.Query("msg")})
}
