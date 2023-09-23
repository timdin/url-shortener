package internal

import (
	"fmt"
	"url-shortener/config"
)

type URLWrapper interface {
	WrapURL(url string) string
}

type URLWrapperImpl struct {
	host string
	port string
}

func NewURLWrapper(config config.ServerConfig) *URLWrapperImpl {
	return &URLWrapperImpl{
		host: config.Host,
		port: config.Port,
	}
}

func (u *URLWrapperImpl) WrapURL(url string) string {
	return fmt.Sprintf("%s:%s/%s", u.host, u.port, url)
}
