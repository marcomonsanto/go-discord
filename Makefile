include .env

build:
	cd backend && go build -o bin/main main.go
	
run:
	cd backend && go run main.go
	
test:
	echo ${DB_STRING}
	
db_status:
	cd backend/db/migrations && goose mysql ${DB_STRING} status

db_up:
	cd backend/db/migrations && goose mysql ${DB_STRING} up
	
db_down:
	cd backend/db/migrations && goose mysql ${DB_STRING} down

db_schema_generate:
	cd backend && sqlc generate