package entity

import (
	"gorm.io/gorm"
)

type (
	Session struct {
		gorm.Model
		Uuid   string
		Email  string
		UserId uint
	}

	Post struct {
		gorm.Model
		Uuid     string `json:"uuid"`
		Body     string `json:"body"`
		UserId   uint   `json:"user_id"`
		ThreadId uint   `json:"thread_id"`
		User     *User
	}

	Thread struct {
		gorm.Model
		Uuid   string
		Topic  string
		UserId uint
		Posts  []*Post
		User   *User
	}

	User struct {
		gorm.Model
		Uuid     string
		Name     string
		Email    string
		Password string
	}
)
