package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/chajiuqqq/chitchat/data"
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&data.FakePost{}))
	writer = httptest.NewRecorder()
}

func TestHandleGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, req)

	if writer.Code != 200 {
		t.Errorf("response code is %d", writer.Code)
	}
	var post data.Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("can't retrieve JSON post")
	}
}
func TestHandlePost(t *testing.T) {
	jsonBody := strings.NewReader(`{"id":5,"uuid":"1239","body":"China ballon9","user_id":1,"thread_id":122,"create_at":"2006-01-02T15:04:05Z07:00"}`)
	req, _ := http.NewRequest("POST", "/post/", jsonBody)
	mux.ServeHTTP(writer, req)

	if writer.Code != 200 {
		t.Errorf("response code is %d,%s", writer.Code, writer.Body)
	}
}
