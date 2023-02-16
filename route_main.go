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

func handleRequest(t data.Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		log.Println(r.Method, r.URL.Path)
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func handleGet(w http.ResponseWriter, r *http.Request, t data.Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	err = t.Fetch(id)
	if err != nil {
		return err
	}
	bytes, err := json.MarshalIndent(t, "", "\t")

	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
	return
}
func handlePost(w http.ResponseWriter, r *http.Request, t data.Text) (err error) {
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return
	}
	err = t.Create()
	return
}
func handlePut(w http.ResponseWriter, r *http.Request, t data.Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = t.Fetch(id)
	if err != nil {
		return
	}
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return
	}
	err = t.Update()
	return
}
func handleDelete(w http.ResponseWriter, r *http.Request, t data.Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = t.Fetch(id)
	if err != nil {
		return
	}
	err = t.Delete()
	if err != nil {
		return
	}
	return
}
