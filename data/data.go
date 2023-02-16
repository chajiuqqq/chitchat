package data

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	Db *sql.DB
)

type (
	Session struct {
		Id       int
		Uuid     string
		Email    string
		UserId   int
		CreateAt time.Time
	}
)

type Text interface {
	Fetch(id int) (err error)
	Create() (err error)
	Update() (err error)
	Delete() (err error)
}

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres password=mkQ445683 dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func UserByEmail(email string) {

}

func (sess *Session) Check() (ok bool, err error) {
	return
}
