package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorMsg(c *gin.Context, msg string) {
	url := []string{"/err?msg=", msg}
	c.Redirect(http.StatusFound, strings.Join(url, ""))
}
