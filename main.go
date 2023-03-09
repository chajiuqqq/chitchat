package main

import (
	"net/http"

	_ "github.com/chajiuqqq/chitchat/data"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msgf("Chitchat %s start at %s", Config.Version, Config.Port)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)

	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    Config.Port,
		Handler: mux,
	}
	server.ListenAndServe()
}
