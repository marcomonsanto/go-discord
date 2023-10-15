include .env.development

build:
	go build -o bin/main main.go
	
run:
	go run main.go

build-dev:
	docker build -f Dockerfile.dev -t go-discord-be:dev .

run-dev:
	docker-compose -f ./docker-compose.dev.yaml up -d

test:
	echo ${DB_STRING}
	
db_status:
	cd db/migrations && goose mysql ${DB_STRING} status

db_up:
	cd db/migrations && goose mysql ${DB_STRING} up
	
db_down:
	cd db/migrations && goose mysql ${DB_STRING} down

db_schema_generate:
	sqlc generate