package util

import (
	"math/rand"
	"strconv"

	_ "github.com/chajiuqqq/chitchat/common/pb"
)

func GenerateUuid() string {
	return strconv.Itoa(rand.Intn(9999))
}
