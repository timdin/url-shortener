package dao

import (
	"context"
	"time"
	"url-shortener/constants"
	"url-shortener/logging"
	"url-shortener/model"

	redis "github.com/go-redis/redis/v8"
)

type CacheInterface interface {
	WriteCache(entity *model.URL) error
	QueryCache(shortURL string) (*model.URL, error)
}

type CacheImpl struct {
	redis *redis.Client
}

func initCache(cacheConfig string) *CacheImpl {
	redisClient := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    cacheConfig,
	})
	return &CacheImpl{
		redis: redisClient,
	}
}

func (cache *CacheImpl) WriteCache(entity *model.URL) error {
	duration := getExpiration(entity.ExpiresAt)
	// delete cache if duration is negative
	if duration < 0 {
		cache.redis.Del(context.Background(), entity.ShortURL)
		return nil
	}
	cache.redis.Set(context.Background(), entity.ShortURL, entity.LongURL, duration)
	return nil
}

func (cache *CacheImpl) QueryCache(shortURL string) (*model.URL, error) {
	res := &model.URL{}
	if longURL, err := cache.redis.Get(context.Background(), shortURL).Result(); err != nil {
		return nil, err
	} else {
		res.LongURL = longURL
		res.ShortURL = shortURL
	}
	logging.SugarLogger.Infow("cache hit",
		"key", res.ShortURL,
		"val", res.LongURL,
		"timestamp", time.Now().Format(time.RFC3339),
	)
	return res, nil
}

func getExpiration(timestamp *time.Time) time.Duration {
	// For no duration in payload, set to default duration
	if timestamp == nil {
		return constants.DEFAULT_EXPIRATION
	}
	validDuration := timestamp.Sub(time.Now())
	// maximum cache expiration time is 1 hour
	// the cache will be updated if the data was expired in cache and got query again or updated by request
	// the cache will be deleted if the data was expired in database
	if validDuration < 0 {
		return -1 * time.Hour
	} else if validDuration > constants.DEFAULT_EXPIRATION {
		return constants.DEFAULT_EXPIRATION
	}
	return validDuration

}
