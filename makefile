validate:
	openapi-generator validate -i ./quest_board_api.yml

run/server:
	openapi-generator generate -i ./quest_board_api.yml -g go-server -o ./server
	GO111MODULE=off go run server/main.go
