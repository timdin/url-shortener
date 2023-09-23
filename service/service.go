package service

import (
	"log"
	"net/http"
	"url-shortener/convert"
	"url-shortener/dao"
	"url-shortener/internal"
	"url-shortener/model"
	"url-shortener/proto/urlshortener"
	"url-shortener/validator"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// db access will be a property
type URLHandler struct {
	mockDB     map[string]string
	db         *dao.MysqlDao
	Redis      *redis.Client
	valid      validator.Validator
	urlWrapper internal.URLWrapper
}

func NewURLHandler(db *dao.MysqlDao, redis *redis.Client, wrapper internal.URLWrapper, validator validator.Validator) *URLHandler {
	return &URLHandler{
		db:         db,
		Redis:      redis,
		urlWrapper: wrapper,
		valid:      validator,
	}
}

func (u *URLHandler) Redirect(c *gin.Context) {
	// get id from path
	log.Println(c.Param("id"))
	// get url from db
	if entity, err := u.db.QueryURLRecord(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		c.Redirect(http.StatusMovedPermanently, entity.LongURL)
	}
}

func (u *URLHandler) Shortern(c *gin.Context) {
	req := &urlshortener.ShorternRequest{}
	res := &urlshortener.ShorternResponse{}
	entity := &model.URL{}
	if err := c.BindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.valid.Validate(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate short url with hash
	entity = convert.ShortenDto2Entity(req)
	entity.ShortURL = internal.HashURL(entity.LongURL)

	// save to db
	if err := u.db.WriteURLRecord(entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res = convert.Entity2ShortenDto(entity, u.urlWrapper)

	c.JSON(http.StatusOK, res)
	return
}
