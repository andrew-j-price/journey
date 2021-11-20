SHELL := /bin/bash

run: build cli_color cli

build:
	CGO_ENABLED=0 go build -o drive

api:
	./drive -api

cli:
	./drive

cli_color:
	./drive -color


docker: docker_stop docker_build docker_run docker_logs

docker_build:
	docker build -t journey .

docker_run:
	docker run --name journey -it -d -p 8080:8080 journey 

docker_stop:
	docker stop journey || true
	docker rm journey || true

docker_logs:
	docker logs -f journey

curl_hello:
	curl -m 2 127.0.0.1:8080
