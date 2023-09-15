gen:
	protoc --go_out=${GOPATH}/src ./proto/*/*.proto

test:
	go test -v ./...