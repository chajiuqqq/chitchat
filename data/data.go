package data

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
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
		User     User
	}

	Thread struct {
		gorm.Model
		Uuid   string
		Topic  string
		UserId uint
		Posts  []Post
		User   User
	}

	User struct {
		gorm.Model
		Uuid     string
		Name     string
		Email    string
		Password string
	}
)

func init() {
	var err error
	dsn := "user=postgres password=mkQ445683 dbname=chitchat sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Db.AutoMigrate(&Post{}, &Thread{}, &Session{}, &User{})
}
