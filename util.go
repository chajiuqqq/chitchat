package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog"

	"github.com/rs/zerolog/log"

	"github.com/chajiuqqq/chitchat/data"
)

var Config Configuration

type Configuration struct {
	Version string `json:"version"`
	Port    string `json:"port"`
	Log     struct {
		LogFile  string `json:"logFile"`
		LogLevel string `json:"logLevel"`
	} `json:"log"`
}

func init() {
	loadConfig()
	file, err := os.OpenFile(Config.Log.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Error().Err(err).Msg("Fail to open log file!")
	}
	switch Config.Log.LogLevel{
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "TEACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}
	log.Logger = zerolog.New(zerolog.MultiLevelWriter(os.Stdout, zerolog.ConsoleWriter{
		Out:        file,
		TimeFormat: "2006-01-02 15:04:05",
	}))
	
}

func loadConfig() {
	bytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Error().Msg("Fail to open config file")
	}
	err = json.Unmarshal(bytes, &Config)
	if err != nil {
		log.Error().Msg("Fail to read config file")
	}
}
func generateHTML(w http.ResponseWriter, data interface{}, fn ...string) {
	files := []string{}
	for _, file := range fn {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))

	templates.ExecuteTemplate(w, "layout", data)

}
func generateUuid() string {
	return strconv.Itoa(rand.Intn(9999))
}

func SessionCheck(w http.ResponseWriter, r *http.Request) (sess *data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = &data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func errorMsg(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

func newSession(u data.User) (sess data.Session) {
	sess = data.Session{
		Uuid:   generateUuid(),
		UserId: u.ID,
		Email:  u.Email,
	}
	data.Db.Create(&sess)
	return
}
