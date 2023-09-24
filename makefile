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

vegeta:
	python3 ./benchmark/gen.py 100
	vegeta attack -rate=10 -duration=30s -targets="benchmark/post.txt" | vegeta report -output="benchmark/post_report.txt"
	vegeta attack -rate=500 -duration=30s -targets="benchmark/get.txt" -redirects=-1 | vegeta report -output="benchmark/get_report.txt"

cleanup:
	rm ./benchmark/data/*.json
	rm ./benchmark/*.txt