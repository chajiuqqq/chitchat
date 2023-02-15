package data

import "time"

type (
	Thread struct {
		Id       int
		Uuid     string
		Topic    string
		UserId   int
		CreateAt time.Time
	}
)

//获取所有thread
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("select id, uuid,topic,user_id,created_at from threads order by created_at desc")
	if err != nil {
		return
	}
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreateAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}

//返回一个thread下的回帖数量
func (t *Thread) NumReplies() (count int) {
	rows, err := Db.Query("select count(*) from posts where thread_id=$1", t.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}
