package validator

import (
	"testing"
	"time"
	"url-shortener/proto/urlshortener"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUrlValidator_Validate(t *testing.T) {
	pastTime := "2020-09-30T00:00:00Z"
	futureTime := "2099-09-30T00:00:00Z"
	noExpiration := ""
	strictValidator := NewUrlValidator(false, false)
	expiredValidator := NewUrlValidator(true, false)
	noExpirationValidator := NewUrlValidator(false, true)
	noLimitValidator := NewUrlValidator(true, true)
	tests := []struct {
		name      string
		validator *UrlValidator
		req       *urlshortener.ShorternRequest
		wantErr   bool
	}{
		{
			name:      "strict validator invalid time format",
			validator: strictValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: "invalid time format",
			},
			wantErr: true,
		},
		{
			name:      "strict validator valid url",
			validator: strictValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: futureTime,
			},
			wantErr: false,
		},
		{
			name:      "strict validator invalid url",
			validator: strictValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "invalid url",
				Expiration: futureTime,
			},
			wantErr: true,
		},
		{
			name:      "strict validator no expiration",
			validator: strictValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: noExpiration,
			},
			wantErr: true,
		},
		{
			name:      "strict validator expired",
			validator: strictValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: pastTime,
			},
			wantErr: true,
		},
		{
			name:      "expired validator invalid time format",
			validator: expiredValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: "invalid time format",
			},
			wantErr: true,
		},
		{
			name:      "expired validator valid url",
			validator: expiredValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: futureTime,
			},
			wantErr: false,
		},
		{
			name:      "expired validator invalid url",
			validator: expiredValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "invalid url",
				Expiration: futureTime,
			},
			wantErr: true,
		},
		{
			name:      "expired validator no expiration",
			validator: expiredValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: noExpiration,
			},
			wantErr: true,
		},
		{
			name:      "expired validator expired",
			validator: expiredValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: pastTime,
			},
			wantErr: false,
		},
		{
			name:      "no expiration validator invalid time format",
			validator: strictValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: "invalid time format",
			},
			wantErr: true,
		},
		{
			name:      "no expiration validator valid url",
			validator: noExpirationValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: futureTime,
			},
			wantErr: false,
		},
		{
			name:      "no expiration validator invalid url",
			validator: noExpirationValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "invalid url",
				Expiration: futureTime,
			},
			wantErr: true,
		},
		{
			name:      "no expiration validator no expiration",
			validator: noExpirationValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: noExpiration,
			},
			wantErr: false,
		},
		{
			name:      "no expiration validator expired",
			validator: noExpirationValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: pastTime,
			},
			wantErr: true,
		},
		{
			name:      "no expiration invalid time format",
			validator: strictValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: "invalid time format",
			},
			wantErr: true,
		},
		{
			name:      "no limit validator valid url",
			validator: noLimitValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: futureTime,
			},
			wantErr: false,
		},
		{
			name:      "no limit validator invalid url",
			validator: noLimitValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "invalid url",
				Expiration: futureTime,
			},
			wantErr: true,
		},
		{
			name:      "no limit validator no expiration",
			validator: noLimitValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: noExpiration,
			},
			wantErr: false,
		},
		{
			name:      "no limit validator expired",
			validator: noLimitValidator,
			req: &urlshortener.ShorternRequest{
				URL:        "https://www.google.com",
				Expiration: pastTime,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.validator.Validate(tt.req); (err != nil) != tt.wantErr {
				t.Errorf("UrlValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func makeTimestamp(timeString string) *timestamppb.Timestamp {
	if t, err := time.Parse("2006-01-02 15:04:05", timeString); err != nil {
		return nil
	} else {
		return timestamppb.New(t)
	}
}
