package main

import (
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/chajiuqqq/chitchat/data"
)

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	files := []string{}
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(w, "layout", data)

}
func generateUuid() string{
	return strconv.Itoa(rand.Intn(9999)) 
}

func SessionCheck(w http.ResponseWriter,r *http.Request)( sess *data.Session,err error){
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = &data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func errorMsg(w http.ResponseWriter,r *http.Request,msg string){
	url := []string{"/err?msg=",msg}
	http.Redirect(w,r,strings.Join(url,""),302)
}

func newSession(u data.User)(sess data.Session){
	sess = data.Session{
		Uuid: generateUuid(),
		UserId: u.ID,
		Email: u.Email,
	}
	data.Db.Create(&sess)
	return
}