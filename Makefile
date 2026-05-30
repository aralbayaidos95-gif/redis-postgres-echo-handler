include .env
export

service-run:
	go run main.go

migrate-up:
	migrate -path db/migrations -database ${CONN_STR} up

migrate-down:
	migrate -path db/migrations -database ${CONN_STR} down
