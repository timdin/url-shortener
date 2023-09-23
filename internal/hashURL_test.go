package internal

import (
	"testing"
)

func TestHashURL(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "test 1",
			url:  "https://www.google.com",
			want: "ac6bb669",
		},
		{
			name: "test 1 - repeat",
			url:  "https://www.google.com",
			want: "ac6bb669",
		},
		{
			name: "test 2",
			url:  "https://www.google.com/search?q=hash+url+golang&oq=hash+url+golang&aqs=chrome..69i57j0l7.2951j0j7&sourceid=chrome&ie=UTF-8",
			want: "d0b59df1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashURL(tt.url); got != tt.want {
				t.Errorf("HashURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
