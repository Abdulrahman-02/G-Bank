postgres:
	docker run --name postgres-1 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
createdb:
	docker exec -it postgres-1 createdb --username=root --owner=root G-Bank

dropdb:
	docker exec -it postgres-1 dropdb G-Bank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/G-Bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/G-Bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server