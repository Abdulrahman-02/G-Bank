postgres:
	docker run --name postgres-1 --network g-bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
createdb:
	docker exec -it postgres-1 createdb --username=root --owner=root G-Bank

dropdb:
	docker exec -it postgres-1 dropdb G-Bank

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/G-Bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/G-Bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/G-Bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/G-Bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/abdulrahman-02/G-Bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup  migratedown migrateup1 migratedown1 sqlc test server mock