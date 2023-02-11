package main

// import (
// 	"net/http"

// 	"github.com/chajiuqqq/chitchat/data"
// )

// func authenticate(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	user, _ := data.UserByEmail(r.Form["email"])
// 	if user.Password == data.Encrypt(r.Form["password"]) {
// 		session := user.NewSession()
// 		cookie := http.Cookie{
// 			Name:     "_cookie",
// 			Value:    session.Uuid,
// 			HttpOnly: true,
// 		}
// 		http.SetCookie(w, &cookie)
// 		http.Redirect(w, r, "/", 302)
// 	} else {
// 		http.Redirect(w, r, "/login", 302)
// 	}
// }
