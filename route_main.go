package main

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/chajiuqqq/chitchat/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		log.Fatalln("can't get threads:", err)
	} else {
		_, err = checkSession(w, r)
		if err == nil {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	log.Println(r.Method,r.URL.Path)
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	post, err := data.GetPost(id)
	if err != nil {
		return err
	}
	bytes, err := json.MarshalIndent(post, "", "\t")

	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	return
}
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	var post data.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		return
	}
	err = post.Create()
	return
}
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := data.GetPost(id)
	if err != nil {
		return
	}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		return
	}
	err = post.Update()
	return
}
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := data.GetPost(id)
	if err != nil {
		return
	}
	err = post.Delete()
	if err != nil {
		return
	}
	return
}
