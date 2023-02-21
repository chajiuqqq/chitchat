package main

import (
	"log"
	"net/http"

	"github.com/chajiuqqq/chitchat/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		log.Fatalln("can't get threads:", err)
	} else {
		_, err = SessionCheck(w, r)
		if err == nil {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		}
	}

}
