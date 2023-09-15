package server

import (
	"fmt"
	"testing"

	"url-shortener/proto/helloworld"

	"google.golang.org/protobuf/proto"
)

func Test(t *testing.T) {
	hello := &helloworld.HelloRequest{
		Name: "Tim",
	}
	data, _ := proto.Marshal(hello)
	res := &helloworld.HelloRequest{}
	err := proto.Unmarshal(data, res)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", res)
}
