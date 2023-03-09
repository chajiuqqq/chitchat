package data

import (
	"encoding/json"
	"log"
	"testing"
)

func TestThreads(t *testing.T) {
	threads, err := Threads()
	if err != nil {
		t.Error(err)
	}
	for _, t := range threads {

		for _, p := range t.Posts {
			log.Println(t.Topic, ":", p.User.Name)
		}
	}
}

func TestGetThread(t *testing.T) {
	th, err := GetThread("713")
	if err != nil {
		t.Fatal(err)
	}
	bytes, _ := json.MarshalIndent(th, "", "\t")
	t.Log(string(bytes))
}
