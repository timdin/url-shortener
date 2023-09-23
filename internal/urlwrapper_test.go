package internal

import "testing"

func TestURLWrapperImpl_WrapURL(t *testing.T) {
	type fields struct {
		host string
		port string
	}
	type args struct {
		url string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "test 1",
			fields: fields{
				host: "http://localhost",
				port: "8080",
			},
			args: args{
				url: "shorturl",
			},
			want: "http://localhost:8080/shorturl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &URLWrapperImpl{
				host: tt.fields.host,
				port: tt.fields.port,
			}
			if got := u.WrapURL(tt.args.url); got != tt.want {
				t.Errorf("URLWrapperImpl.WrapURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
