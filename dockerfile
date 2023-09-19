FROM golang:latest

RUN mkdir -p /usr/local/go/src/url-shortener
WORKDIR /usr/local/go/src/url-shortener
ADD . /usr/local/go/src/url-shortener

RUN go mod download
RUN go build ./main.go

EXPOSE 8080

CMD ["./main"]