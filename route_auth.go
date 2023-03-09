package main

import (
	"net/http"

	"github.com/chajiuqqq/chitchat/data"
	"github.com/gin-gonic/gin"
)

func authenticate(c *gin.Context) {
	user, _ := data.UserByEmail(c.PostForm("email"))
	if user.Password == data.Encrypt(c.PostForm("password")) {
		session := newSession(user)
		c.SetCookie("_cookie", session.Uuid, 3600, "/", "localhost", true, true)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.Redirect(http.StatusFound, "/login")
	}

}

func login(c *gin.Context) {
	c.HTML(200, "login.tmpl", gin.H{"IsPublic": true})
}

func logout(c *gin.Context) {
	c.SetCookie("_cookie", "", -1, "/", "localhost", true, true)
	c.Redirect(http.StatusFound, "/")
}

func signup(c *gin.Context) {
	c.HTML(200, "signup.tmpl", gin.H{"IsPublic": true})
}

func signupAccount(c *gin.Context) {
	user := data.User{
		Name:     c.PostForm("name"),
		Email:    c.PostForm("email"),
		Password: data.Encrypt(c.PostForm("password")),
		Uuid:     generateUuid(),
	}
	data.Db.Create(&user)
	c.Redirect(http.StatusFound, "/")
}

func err(c *gin.Context) {
	_, err := SessionCheck(c.Writer, c.Request)
	c.HTML(200, "error.tmpl", gin.H{"IsPublic": err != nil, "Msg": c.Query("msg")})
}
