package data

import "time"

type (
	Post struct {
		Id       int
		Uuid     string
		Body     string
		UserId   int
		ThreadId int
		CreateAt time.Time
	}
)

func Posts(limits int) (posts []Post, err error) {
	rows, err := Db.Query("select * from posts limit $1",limits)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreateAt)
		if err!=nil{
			return nil,err
		}
		posts = append(posts, post)
	}
	return
}

func GetPost(id int)(post Post,err error){
	post = Post{}
	err = Db.QueryRow("select * from posts where id=$1",id).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreateAt)
	if err!=nil{
		return Post{},err
	}
	return
}

func (post *Post) Create() (err error){
	res,err := Db.Exec("insert into posts values($1,$2,$3,$4,$5,$6)",post.Id,post.Uuid,post.Body,post.UserId,post.ThreadId,post.CreateAt)
	if err!=nil{
		return err
	}
	id,err:= res.LastInsertId()
	if err!=nil{
		return err
	}
	post.Id = int(id)
	return
}

//全量更新
func (post *Post) Update() (err error){
	_,err = Db.Exec("update posts set uuid=$1,body=$2,user_id=$3,thread_id=$4 where id=$5",post.Uuid,post.Body,post.UserId,post.ThreadId,post.Id)
	return
}

func (post *Post) Delete()(err error){
	_,err = Db.Exec("delete from posts where id=$1",post.Id)
	return
}