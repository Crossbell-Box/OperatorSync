all: build
docker: build-docker
dev: dev-start
prod: prod-start

build-docker:
	docker build . -f deploy/dockerfile/server.Dockerfile -t rss3/crossbell-operator-sync:server
	docker build . -f deploy/dockerfile/worker.Dockerfile -t rss3/crossbell-operator-sync:worker

dev-start:
	docker-compose -f deploy/docker-compose/docker-compose.dev.yml -p operatorsync-dev up -d

dev-stop:
	docker-compose -f deploy/docker-compose/docker-compose.dev.yml stop

dev-down:
	docker-compose -f deploy/docker-compose/docker-compose.dev.yml down

prod-start: build-docker
	docker-compose -f deploy/docker-compose/docker-compose.prod.yml -p operatorsync up -d

prod-stop:
	docker-compose -f deploy/docker-compose/docker-compose.prod.yml stop

prod-down:
	docker-compose -f deploy/docker-compose/docker-compose.prod.yml down

clean-docker: dev-down prod-down
	docker image rm rss3/crossbell-operator-sync:server -f
	docker image rm rss3/crossbell-operator-sync:worker -f

.PHONY : build
build:
	go build -o ./build/ ./app/server
	go build -o ./build/ ./app/worker

clean:
	rm -rf ./build/*
