package main

import (
	"fmt"
	"net/http"

	"github.com/chajiuqqq/chitchat/data"
)

func newThread(w http.ResponseWriter, r *http.Request) {
	_, err := SessionCheck(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}

	generateHTML(w, nil, "layout", "new.thread", "private.navbar")
}

func createThread(w http.ResponseWriter, r *http.Request) {
	sess, err := SessionCheck(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	thread := data.Thread{
		Topic:  r.FormValue("topic"),
		Uuid:   generateUuid(),
		UserId: sess.UserId,
	}
	data.Db.Create(&thread)
	http.Redirect(w, r, "/", 302)

}

//对一个thread发表post
func postThread(w http.ResponseWriter, r *http.Request) {
	sess, err := SessionCheck(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	uuid := r.FormValue("uuid")
	body := r.FormValue("body")
	var thread data.Thread
	data.Db.Where("uuid=?", uuid).First(&thread)
	data.Db.Model(&thread).Association("Posts").Append(
		&data.Post{Uuid: generateUuid(), Body: body, UserId: sess.UserId},
	)

	url := fmt.Sprint("/thread/read?id=", uuid)
	http.Redirect(w, r, url, 302)
}

func readThread(w http.ResponseWriter, r *http.Request) {
	uuid := r.FormValue("id")
	thread := data.Thread{}
	err := data.Db.Preload("Posts").Where("uuid=?", uuid).First(&thread).Error
	if err != nil {
		errorMsg(w, r, err.Error())
	}
	if _, err = SessionCheck(w, r); err == nil {
		generateHTML(w, &thread, "layout", "private.navbar", "private.thread")
	} else {
		generateHTML(w, &thread, "layout", "public.navbar", "public.thread")
	}
}
