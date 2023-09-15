package server

import (
	"fmt"
	"testing"

	"url-shortener/proto/helloworld"

	"google.golang.org/protobuf/proto"
)

func Test(t *testing.T) {
	fmt.Println("test")
	hello := &helloworld.HelloRequest{
		Name: "Tim",
	}
	data, _ := proto.Marshal(hello)
	fmt.Println(data)
}
