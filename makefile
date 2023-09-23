gen:
	protoc --go_out=${GOPATH}/src ./proto/*/*.proto

test:
	go test -v ./...

run:
	docker compose up

build:
	docker compose build

update:
	make build
	make run