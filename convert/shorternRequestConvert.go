package convert

import (
	"time"
	"url-shortener/constants"
	"url-shortener/internal"
	"url-shortener/model"
	"url-shortener/proto/urlshortener"
)

// ShortenDto2Entity: convert ShorternRequest to URL entity
// NOTE: since the protobuf timestamp has poor json support, we use string instead
func ShortenDto2Entity(dto *urlshortener.ShorternRequest) *model.URL {
	// since the time format has been validated in validator, we can safely parse it here
	expTimestamp, _ := time.Parse(constants.TIME_FORMAT, dto.GetExpiration())
	expTime := &expTimestamp
	if expTimestamp.IsZero() {
		expTime = nil
	}
	res := &model.URL{
		LongURL:   dto.GetURL(),
		ExpiresAt: expTime,
	}
	return res
}

func Entity2ShortenDto(entity *model.URL, wrapper internal.URLWrapper) *urlshortener.ShorternResponse {
	res := &urlshortener.ShorternResponse{
		ShortURL: wrapper.WrapURL(entity.ShortURL),
		ID:       entity.ShortURL,
	}
	return res
}
