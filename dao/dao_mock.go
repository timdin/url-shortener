package dao

import "url-shortener/model"

type DaoMock struct {
	mockStorage map[string]*model.URL
}

func NewDaoMock() *DaoMock {
	return &DaoMock{
		mockStorage: make(map[string]*model.URL),
	}
}

func (d *DaoMock) WriteURLRecord(entity *model.URL) error {
	d.mockStorage[entity.ShortURL] = entity
	return nil
}

func (d *DaoMock) QueryURLRecord(shortURL string) (*model.URL, error) {
	if entity, ok := d.mockStorage[shortURL]; ok {
		return entity, nil
	}
	return nil, nil
}
