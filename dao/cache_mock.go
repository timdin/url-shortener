package dao

import (
	"errors"
	"fmt"
	"url-shortener/model"
)

type cacheMock struct {
	mockStorage map[string]string
}

func newCacheMock() *cacheMock {
	return &cacheMock{
		mockStorage: make(map[string]string),
	}
}

func (c *cacheMock) WriteCache(entity *model.URL) error {
	c.mockStorage[entity.ShortURL] = entity.LongURL
	return nil
}

func (c *cacheMock) QueryCache(shortURL string) (*model.URL, error) {
	fmt.Println("query cache")
	if longURL, ok := c.mockStorage[shortURL]; ok {
		return &model.URL{
			ShortURL: shortURL,
			LongURL:  longURL,
		}, nil
	}
	fmt.Println("cache miss")
	return nil, errors.New("cache miss")
}
