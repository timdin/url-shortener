package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// db access will be a property
type URLHandler struct {
	mockDB map[string]string
}

func NewURLHandler() *URLHandler {
	return &URLHandler{}
}

func NewMockURLHandler() *URLHandler {
	return &URLHandler{
		mockDB: map[string]string{
			"123": "https://www.google.com/",
			"456": "https://www.yahoo.com/",
		},
	}
}

func (u *URLHandler) Redirect(c *gin.Context) {
	id := c.Param("id")
	if v, ok := u.mockDB[id]; ok {
		c.Redirect(http.StatusMovedPermanently, v)
	} else {
		c.Errors = append(c.Errors, &gin.Error{
			Err: fmt.Errorf("id %s not found", id),
		})
	}
}

func (u *URLHandler) Shortern(c *gin.Context) {
	// req := urlshortener.ShorternRequest{}

}
