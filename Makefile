SHELL := /bin/bash

run: build cli_color cli_add cli_random

build:
	CGO_ENABLED=0 go build -o drive

cli:
	./drive

cli_color:
	./drive -debug -color

cli_add:
	./drive -math add 5 7

cli_random:
	./drive -random


# testing
test: unit_test

unit_test:
	go test -v -cover


# NOTE: analyze golangci-lint
fmt:
	go fmt


# docker workflow
docker: docker_build docker_compose_build docker_compose_logs
down: docker_stop docker_compose_down

docker_compose_build:
	docker-compose build && \
	docker-compose --profile core up -d

docker_compose_logs:
	docker-compose logs -f drive

docker_compose_exec:
	docker-compose exec drive sh

docker_compose_exec_scratch:
	docker-compose exec drive ./drive -debug -color

docker_compose_down:
	docker-compose down --remove-orphans

docker_build:
	docker build -t journey .

docker_run:
	docker run --name journey -it -d --env LISTEN_ADDRESS=":8080" -p 8080:8080 journey 

docker_stop:
	docker stop journey || true
	docker rm journey || true

docker_logs:
	docker logs -f journey

docker_exec:
	docker exec -it journey sh

docker_cli_color:
	docker exec -t journey sh -c "./drive -color"


# helper calls
curl_hello:
	curl -m 2 127.0.0.1:8080


# prefer docker method over these api methods
api_start:
	bash -c "./drive -api &" && \
	echo ""

api_stop:
	kill `lsof -ti:8080`


# cleanup items
cleanup: delete_binaries delete_logs

delete_binaries:
	rm -f ./drive && \
	rm -f ./journey

delete_logs:
	rm -f ./log.txt