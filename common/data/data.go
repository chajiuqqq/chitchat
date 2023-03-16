package data

import (
	"log"

	. "github.com/chajiuqqq/chitchat/common/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func init() {
	var err error
	dsn := "host=db user=chitchat password=chitchatpass dbname=chitchat sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Db.AutoMigrate(&Post{}, &Thread{}, &Session{}, &User{})
}
