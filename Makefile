CURRENT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))


up:
	make containerup
	make createdb
	make migrationup
	
containerup:
	docker run --name backend-masterclass-db --network bank-network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:15

containerdown:
	docker stop backend-masterclass-db
	docker rm --force backend-masterclass-db

createdb: 
	docker exec -it backend-masterclass-db createdb --username=root --owner=root bank

dropdb:
	docker exec -it backend-masterclass-db dropdb bank

migrationup:
	migrate -path db/migrations -database "postgres://root:mysecretpassword@localhost:5432/bank?sslmode=disable" --verbose up

migrationdown:
	migrate -path db/migrations -database "postgres://root:mysecretpassword@localhost:5432/bank?sslmode=disable" --verbose down 1

sqlc: 
	docker run --rm -v ${CURRENT_DIR}:/src -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

server: 
	go run main.go

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: postgres createdb dropdb migrationup migrationdown sqlc test respawn proto