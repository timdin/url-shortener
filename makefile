gen:
	protoc --go_out=${GOPATH}/src ./proto/*/*.proto

test:
	go test -v ./...

run:
	go run main.go

build:
	docker compose build