package dao

import (
	"errors"
	"time"
	"url-shortener/config"
	"url-shortener/constants"
	"url-shortener/model"
)

type DaoImpl struct {
	db    DBinterface
	cache CacheInterface
}

type Dao interface {
	WriteURLRecord(entity *model.URL) error
	QueryURLRecord(shortURL string) (*model.URL, error)
}

func InitDao(config *config.Config) *DaoImpl {
	db, err := initDB(config.DB.Conn)
	if err != nil {
		panic(err)
	}
	cache := initCache(config.Cache.Conn)

	return &DaoImpl{
		db:    db,
		cache: cache,
	}
}

func (dao *DaoImpl) WriteURLRecord(entity *model.URL) error {
	// write db first, then cache
	// fail fast
	if err := dao.db.WriteDB(entity); err != nil {
		return err
	}
	if err := dao.cache.WriteCache(entity); err != nil {
		return err
	}
	return nil

}

func (dao *DaoImpl) QueryURLRecord(shortURL string) (*model.URL, error) {
	// Query cache first, if not found or has error, query db
	// fail slow
	if entity, err := dao.cache.QueryCache(shortURL); err != nil {
		// early return if cache receives empty result
		if entity != nil && entity.LongURL == "" {
			return nil, errors.New("record not found")
		}
		if entity, err := dao.db.QueryDB(shortURL); err != nil {
			// write invalid request to cache
			if constants.CACHE_INVALID_REQUESTS {
				now := time.Now().Add(constants.INVALID_REQUEST_EXPIRATION)
				if constants.CACHE_INVALID_REQUESTS {
					dao.cache.WriteCache(&model.URL{
						ShortURL:  shortURL,
						ExpiresAt: &now,
					})
				}
			}
			return nil, err
		} else {
			// update cache data if found in db
			dao.cache.WriteCache(entity)
			return entity, nil
		}
	} else {
		return entity, nil
	}
}
