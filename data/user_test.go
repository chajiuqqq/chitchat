package data

import (
	"testing"
)

func TestUserByEmail(t *testing.T) {
	u, err := UserByEmail("chajiuqqq@gmail.com")
	if err != nil {
		t.Fatal(err)
	}
	if u.Name != "chajiu" {
		t.Error("name error,", u.Name)
	}
}
