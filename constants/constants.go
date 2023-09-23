package constants

import "time"

const (
	// SHORT_URL_LENGTH: the length of the shortern url
	SHORT_URL_LENGTH = 8
	// time format
	TIME_FORMAT = time.RFC3339
	// default expiration time
	DEFAULT_EXPIRATION = 1 * time.Hour
	// cache invalid requests or not
	CACHE_INVALID_REQUESTS = true
	// cache invalid requests expiration
	INVALID_REQUEST_EXPIRATION = 5 * time.Minute
)
