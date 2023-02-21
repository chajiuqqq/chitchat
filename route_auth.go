package main

import (
	"net/http"
	"time"

	"github.com/chajiuqqq/chitchat/data"
)

func authenticate(w http.ResponseWriter, r *http.Request) {
	user, _ := data.UserByEmail(r.FormValue("email"))
	if user.Password == data.Encrypt(r.FormValue("password")) {
		session := newSession(user)
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "login", "public.navbar")
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "_cookie",
		HttpOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}

func signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "login.layout", "signup", "public.navbar")
}

func signupAccount(w http.ResponseWriter, r *http.Request) {
	user := data.User{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: data.Encrypt(r.FormValue("password")),
		Uuid:     generateUuid(),
	}
	data.Db.Create(&user)
	http.Redirect(w, r, "/login", 302)
}

func err(w http.ResponseWriter, r *http.Request) {
	_, err := SessionCheck(w, r)
	if err == nil {
		generateHTML(w, r.FormValue("msg"), "layout", "private.navbar", "error")
	} else {
		generateHTML(w, r.FormValue("msg"), "layout", "public.navbar", "error")
	}
}
