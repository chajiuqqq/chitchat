package main

import (
	"net/http"

	"github.com/chajiuqqq/chitchat/data"
	"github.com/rs/zerolog/log"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		log.Error().Err(err).Msg("can't get threads")
		errorMsg(w,r,"can't get threads")
	} else {
		_, err = SessionCheck(w, r)
		if err == nil {
			generateHTML(w, threads, "layout", "private.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "public.navbar", "index")
		}
	}

}
