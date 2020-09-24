

build:
	docker build -t myriaddreamin/boj-server:latest -f deployment/docker/Dockerfile .
.PHONY: build

up:
	docker-compose -f deployment/docker/docker-compose.yml up
.PHONY: up

down:
	docker-compose -f deployment/docker/docker-compose.yml down
.PHONY: down

start:
	docker-compose -f deployment/docker/docker-compose.yml start
.PHONY: start

stop:
	docker-compose -f deployment/docker/docker-compose.yml stop
.PHONY: stop



