package validator

import (
	"errors"
	"net/url"
	"time"
	"url-shortener/constants"
	"url-shortener/messages"
	"url-shortener/proto/urlshortener"
)

type Validator interface {
	Validate(req *urlshortener.ShorternRequest) error
}

type UrlValidator struct {
	acceptExpired  bool
	acceptNoExpire bool
}

func NewUrlValidator(acceptExpired bool, acceptNoExpire bool) *UrlValidator {
	return &UrlValidator{
		acceptExpired:  acceptExpired,
		acceptNoExpire: acceptNoExpire,
	}
}

// validate url & request body
// NOTE: this validation is only validating the format of the url, not the existence of the url
func (validator *UrlValidator) Validate(req *urlshortener.ShorternRequest) error {
	if err := validator.validateURL(req); err != nil {
		return err
	}
	if err := validator.validateWithExpiration(req); err != nil {
		return err
	}
	return nil
}

func (validator *UrlValidator) validateURL(req *urlshortener.ShorternRequest) error {
	if _, err := url.ParseRequestURI(req.GetURL()); err != nil {
		return err
	}
	return nil
}

// validate expiration date
func (validator *UrlValidator) validateWithExpiration(req *urlshortener.ShorternRequest) error {
	expString := req.GetExpiration()
	if expString == "" {
		if validator.acceptNoExpire {
			return nil
		}
		return errors.New(messages.ExpirationRequired)
	}
	if expTimestamp, err := time.Parse(constants.TIME_FORMAT, expString); err != nil {
		return errors.New(messages.ExpirationInvalidFormat)
	} else if expTimestamp.Before(time.Now()) {
		if validator.acceptExpired {
			return nil
		}
		return errors.New(messages.ExpirationInPast)
	}
	return nil
}
