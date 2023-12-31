package service

import (
	"log"
	"net/http"
	"time"
	"url-shortener/config"
	"url-shortener/convert"
	"url-shortener/dao"
	"url-shortener/internal"
	"url-shortener/logging"
	"url-shortener/model"
	"url-shortener/proto/urlshortener"
	"url-shortener/validator"

	"github.com/gin-gonic/gin"
)

// db access will be a property
type URLHandler struct {
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
	// get url from storage
	// if not found or invalid cache was hit, return 404
	if entity, err := u.dao.QueryURLRecord(c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	} else {
		log.Println(internal.DumpStruct(entity))
		logging.SugarLogger.Infow("redirect",
			"resp", internal.DumpStruct(entity),
			"timestamp", time.Now().Format(time.RFC3339),
		)
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
	logging.SugarLogger.Infow("shortern",
		"req", internal.DumpStruct(req),
		"timestamp", time.Now().Format(time.RFC3339),
	)

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
	logging.SugarLogger.Infow("shortern",
		"resp", internal.DumpStruct(res),
		"timestamp", time.Now().Format(time.RFC3339),
	)

	c.JSON(http.StatusOK, res)
	return
}
