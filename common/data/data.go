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
	dsn := "user=postgres password=mkQ445683 dbname=chitchat sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Db.AutoMigrate(&Post{}, &Thread{}, &Session{}, &User{})
}
