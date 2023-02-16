package data

import (
	"database/sql"
	"time"
)

type (
	FakePost struct {
		Db       *sql.DB
		Id       int       `json:"id"`
		Uuid     string    `json:"uuid"`
		Body     string    `json:"body"`
		UserId   int       `json:"user_id"`
		ThreadId int       `json:"thread_id"`
		CreateAt time.Time `json:"create_at"`
	}
)

func (post *FakePost) Fetch(id int) (err error) {
	return
}

func (post *FakePost) Create() (err error) {

	return
}

// 全量更新
func (post *FakePost) Update() (err error) {

	return
}

func (post *FakePost) Delete() (err error) {

	return
}
