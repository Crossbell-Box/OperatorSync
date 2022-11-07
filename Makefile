all: build
docker: build-docker
dev: dev-start
prod: prod-start

.PHONY : build
build: build/executable

clean: clean-build

build-docker:
	docker build . -f deploy/dockerfile/server.Dockerfile -t rss3/crossbell-xsync:server
	docker build . -f deploy/dockerfile/worker.Dockerfile -t rss3/crossbell-xsync:worker

build/contract:
	mkdir -p build/contract/
	cd Crossbell-Contracts/ && \
	yarn && \
	make install && \
	make build && \
	forge inspect Web3Entry abi > ../build/contract/Web3Entry.abi && \
	forge inspect Web3Entry b > ../build/contract/Web3Entry.bin && \
	cd ../ && \
	mkdir -p app/worker/chain/contract/
	abigen --bin=build/contract/Web3Entry.bin --abi=build/contract/Web3Entry.abi --pkg=contract --out=app/worker/chain/contract/Web3Entry.go

dev-start:
	docker-compose -f deploy/docker-compose/docker-compose.dev.yml -p operatorsync-dev up -d

dev-stop:
	docker-compose -f deploy/docker-compose/docker-compose.dev.yml stop

dev-down:
	docker-compose -f deploy/docker-compose/docker-compose.dev.yml down

prod-start:
	docker-compose -f deploy/docker-compose/docker-compose.prod.yml -p operatorsync up -d

prod-stop:
	docker-compose -f deploy/docker-compose/docker-compose.prod.yml stop

prod-down:
	docker-compose -f deploy/docker-compose/docker-compose.prod.yml down

clean-docker: dev-down prod-down
	docker image rm rss3/crossbell-xsync:server -f
	docker image rm rss3/crossbell-xsync:worker -f

build/executable:
	go build -o ./build/executable/ ./app/server
	go build -o ./build/executable/ ./app/worker

clean-build:
	rm -rf ./build/*
