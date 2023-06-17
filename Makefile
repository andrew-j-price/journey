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
	go test ./... -v -cover

unit_test_main:
	go test -v -cover


# documentation (only pascal case items are documented)
go_doc_all:
	go doc -all

go_doc_module:
	go doc github.com/andrew-j-price/journey

go_doc_module_stringtoint:
	go doc github.com/andrew-j-price/journey/helpers.StringToInt

go_doc_fuction_stringtoint:
	go doc helpers.StringToInt

# https://pkg.go.dev/fmt#Printf
go_doc_fmt_printf:
	go doc fmt.Printf

# https://pkg.go.dev/net/http#Handler
go_doc_net_http_handler:
	go doc net/http.handler

# run without building
go_run_dir:
	go run . -debug

go_run_module:
	go run github.com/andrew-j-price/journey -debug


# NOTE: analyze golangci-lint
fmt:
	go fmt


# docker workflow
docker: docker_build docker_compose_build docker_compose_logs
down: docker_stop docker_compose_down

docker_compose_build:
	docker compose build && \
	docker compose --profile core up -d

docker_compose_logs:
	docker compose logs -f drive

docker_compose_exec:
	docker compose exec drive sh

docker_compose_exec_scratch:
	docker compose exec drive ./drive -debug -color

docker_compose_down:
	docker compose down --remove-orphans

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
api: api_stop_fuser build api_start

api_start:
	bash -c "./drive -api &" && \
	echo ""

api_stop_lsof:
	kill `lsof -ti:8080`

api_stop_fuser:
	result=`fuser -k -n tcp 8080 || true`

# identity
identity: api_stop_fuser build identity_start

identity_start:
	bash -c "./drive -identity &" && \
	echo ""

identity_docker: identity_docker_build identity_docker_push

identity_docker_build:
	docker build -t andrewprice/identity:v1 . -f Dockerfile-identity

identity_docker_push:
	docker push andrewprice/identity:v1

# cleanup items
cleanup: delete_binaries delete_logs

delete_binaries:
	rm -f ./drive && \
	rm -f ./journey

delete_logs:
	rm -f ./log.txt
