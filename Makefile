all: build-image
dev: dev-start
prod: prod-start

build-image:
	docker build . -f deploy/dockerfile/server.Dockerfile -t operator-sync/server
	docker build . -f deploy/dockerfile/worker.Dockerfile -t operator-sync/worker

dev-start:
	docker-compose -f deploy/docker-compose/docker-compose.dev.yml -p operatorsync-dev up -d

dev-stop:
	docker-compose -f deploy/docker-compose/docker-compose.dev.yml stop

dev-down:
	docker-compose -f deploy/docker-compose/docker-compose.dev.yml down

prod-start: build-image
	docker-compose -f deploy/docker-compose/docker-compose.prod.yml -p operatorsync up -d

prod-stop:
	docker-compose -f deploy/docker-compose/docker-compose.prod.yml stop

prod-down:
	docker-compose -f deploy/docker-compose/docker-compose.prod.yml down

clean: dev-down prod-down
	docker image rm operator-sync/server -f
	docker image rm operator-sync/worker -f
