package service

import (
	"log"
	"net/http"
	"url-shortener/config"
	"url-shortener/convert"
	"url-shortener/dao"
	"url-shortener/internal"
	"url-shortener/model"
	"url-shortener/proto/urlshortener"
	"url-shortener/validator"

	"github.com/gin-gonic/gin"
)

// db access will be a property
type URLHandler struct {
	mockDB     map[string]string
	dao        dao.Dao
	valid      validator.Validator
	urlWrapper internal.URLWrapper
}

func NewURLHandler(config *config.Config) *URLHandler {
	return &URLHandler{
		dao:        dao.InitDao(config),
		urlWrapper: internal.NewURLWrapper(config.Server),
		valid:      validator.NewUrlValidator(config.Server.AcceptExpired, config.Server.AcceptNoExpire),
	}
}

func (u *URLHandler) Redirect(c *gin.Context) {
	// get id from path
	log.Println(c.Param("id"))
	// get url from storage
	// if not found or invalid cache was hit, return 404
	if entity, err := u.dao.QueryURLRecord(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		log.Println(internal.DumpStruct(entity))
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
	if err := u.dao.WriteURLRecord(entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res = convert.Entity2ShortenDto(entity, u.urlWrapper)

	c.JSON(http.StatusOK, res)
	return
}
