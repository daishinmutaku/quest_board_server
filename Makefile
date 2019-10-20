DB_REPOSITORY_NAME:=quest-board/db
DB_CONTAINER_NAME:=quest-board-db-dev

SERVER_REPOSITORY_NAME:=quest-board/server
SERVER_CONTAINER_NAME:=quest-board-server-dev
HOST_SERVER_BASE:=$(shell pwd)/server
DOCKER_SERVER_BASE:=/go/src/github.com/daishinmutaku/quest_board_server/server

DB_VOLUME_PATH:=$(shell pwd)/_secret/quest-board-data
DB_DATA_PATH:=/var/lib/mysql

DBNAME:=quest-board
ENV:=development

validate:
	openapi-generator validate -i ./quest_board_api.yml

run/mock:
	openapi-generator generate -i ./quest_board_api.yml -g go-server -o ./mock-server
	GO111MODULE=off go run mock-server/main.go

docker/build/db:
	docker build -t $(DB_REPOSITORY_NAME) ./db

docker/build/server:
	cd ./server && docker build -t $(SERVER_REPOSITORY_NAME) .

docker/run/db:
	docker run -d -p 3306:3306 --name $(DB_CONTAINER_NAME) -v $(DB_VOLUME_PATH):$(DB_DATA_PATH) --env-file _secret/.env $(DB_REPOSITORY_NAME):latest
	@echo 'Connect DB port :3306!!!'

docker/run/server:
	docker run -p 8080:8080 --name $(SERVER_CONTAINER_NAME) --env-file _secret/.env -v $(HOST_SERVER_BASE):$(DOCKER_SERVER_BASE) $(SERVER_REPOSITORY_NAME):latest

docker/run/server/d:
	docker run -d -p 8080:8080 --name $(SERVER_CONTAINER_NAME) --env-file _secret/.env -v $(HOST_SERVER_BASE):$(DOCKER_SERVER_BASE) $(SERVER_REPOSITORY_NAME):latest

docker/stop/db:
	docker container stop $(DB_CONTAINER_NAME)
	docker container rm $(DB_CONTAINER_NAME)

docker/stop/server:
	docker container stop $(SERVER_CONTAINER_NAME)
	docker container rm $(SERVER_CONTAINER_NAME)

migrate/up:
	sql-migrate up

migrate/down:
	sql-migrate down

migrate/status:
	sql-migrate status
