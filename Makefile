all: build
docker: build-docker
dev: dev-start
prod: prod-start

.PHONY : build
build: build/contract build/executable

clean: clean-build clean-docker

build-docker:
	docker build . -f deploy/dockerfile/server.Dockerfile -t rss3/crossbell-operator-sync:server
	docker build . -f deploy/dockerfile/worker.Dockerfile -t rss3/crossbell-operator-sync:worker

build/contract:
	solc --abi Crossbell-Contracts/src/Web3Entry.sol > build/contract/Web3Entry.abi
	solc --bin Crossbell-Contracts/src/Web3Entry.sol > build/contract/Web3Entry.bin
	abigen --bin=build/contract/Web3Entry.bin --abi=build/contract/Web3Entry.abi --pkg=contract --out=worker/chain/contract/Web3Entry.go

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

build/executable:
	go build -o ./build/executable/ ./app/server
	go build -o ./build/executable/ ./app/worker

clean-build:
	rm -rf ./build/*
