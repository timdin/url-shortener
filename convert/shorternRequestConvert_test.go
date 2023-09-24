package convert

import (
	"reflect"
	"testing"
	"time"
	"url-shortener/config"
	"url-shortener/constants"
	"url-shortener/internal"
	"url-shortener/model"
	"url-shortener/proto/urlshortener"
)

func TestShortenDto2Entity(t *testing.T) {
	type args struct {
		dto *urlshortener.ShorternRequest
	}
	timeString := "2021-01-01T00:00:00Z"
	timeStamp, _ := time.Parse(constants.TIME_FORMAT, timeString)
	tests := []struct {
		name string
		args args
		want *model.URL
	}{
		{
			name: "with timestamp",
			args: args{
				dto: &urlshortener.ShorternRequest{
					URL:        "https://www.google.com",
					Expiration: "2021-01-01T00:00:00Z",
				},
			},
			want: &model.URL{
				LongURL:   "https://www.google.com",
				ExpiresAt: &timeStamp,
			},
		},
		{
			name: "with out timestamp",
			args: args{
				dto: &urlshortener.ShorternRequest{
					URL:        "https://www.google.com",
					Expiration: "",
				},
			},
			want: &model.URL{
				LongURL:   "https://www.google.com",
				ExpiresAt: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortenDto2Entity(tt.args.dto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortenDto2Entity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity2ShortenDto(t *testing.T) {
	type args struct {
		entity  *model.URL
		wrapper internal.URLWrapper
	}
	urlWrapper1 := internal.NewURLWrapper(config.ServerConfig{
		Host: "http://localhost",
		Port: "8080",
	})
	urlWrapper2 := internal.NewURLWrapper(config.ServerConfig{
		Host: "http://otherhost",
		Port: "8080",
	})
	tests := []struct {
		name    string
		wrapper internal.URLWrapper
		entity  *model.URL
		want    *urlshortener.ShorternResponse
	}{
		{
			name:    "wrapper1",
			wrapper: urlWrapper1,
			entity: &model.URL{
				ShortURL: "abc",
			},
			want: &urlshortener.ShorternResponse{
				ShortURL: "http://localhost:8080/abc",
				ID:       "abc",
			},
		},
		{
			name:    "wrapper2",
			wrapper: urlWrapper2,
			entity: &model.URL{
				ShortURL: "abc",
			},
			want: &urlshortener.ShorternResponse{
				ShortURL: "http://otherhost:8080/abc",
				ID:       "abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Entity2ShortenDto(tt.entity, tt.wrapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Entity2ShortenDto() = %v, want %v", got, tt.want)
			}
		})
	}
}
