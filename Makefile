DB_REPOSITORY_NAME:=quest-board/db
DB_CONTAINER_NAME:=quest-board-db-dev

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

docker/run/db:
	docker run -d -p 3306:3306 --name $(DB_CONTAINER_NAME) -v $(DB_VOLUME_PATH):$(DB_DATA_PATH) --env-file _secret/.env $(DB_REPOSITORY_NAME):latest
	@echo 'Connect DB port :3306!!!'

docker/stop/db:
	docker container stop $(DB_CONTAINER_NAME)
	docker container rm $(DB_CONTAINER_NAME)

migrate/up:
	sql-migrate up

migrate/down:
	sql-migrate down

migrate/status:
	sql-migrate status
