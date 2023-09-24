package dao

import (
	"errors"
	"fmt"
	"url-shortener/model"
)

type databaseMock struct {
	mockStorage map[string]*model.URL
}

func newDatabaseMock() *databaseMock {
	return &databaseMock{
		mockStorage: make(map[string]*model.URL),
	}
}

func (d *databaseMock) WriteDB(entity *model.URL) error {
	d.mockStorage[entity.ShortURL] = entity
	return nil
}

func (d *databaseMock) QueryDB(shortURL string) (*model.URL, error) {
	fmt.Println("query db")
	if entity, ok := d.mockStorage[shortURL]; ok {
		return entity, nil
	}
	return nil, errors.New("db miss")
}
