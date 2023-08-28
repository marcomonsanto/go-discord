include .env

build:
	go build -o bin/main main.go
	
run:
	go run main.go
	
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