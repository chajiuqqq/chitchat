package data

import (
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
