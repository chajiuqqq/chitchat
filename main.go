package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	files := []string{}
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(w, "layout", data)

}
func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/post/", handleRequest)
	// mux.HandleFunc("/err", err)
	// mux.HandleFunc("/login", login)
	// mux.HandleFunc("/logout", logout)
	// mux.HandleFunc("/signup", signup)
	// mux.HandleFunc("/signup_account", signupAccount)
	// mux.HandleFunc("/authenticate", authenticate)
	// mux.HandleFunc("/thread/new", newThread)
	// mux.HandleFunc("/thread/create", createThread)
	// mux.HandleFunc("/thread/post", postThread)
	// mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("start server at 8080")
	server.ListenAndServe()
}
