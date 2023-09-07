DB_URL = postgresql://root:password@localhost:5432/G-Bank?sslmode=disable

postgres:
	docker run --name postgres-1 --network g-bank-network -p 5432:5432 \
	-e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
createdb:
	docker exec -it postgres-1 createdb --username=root --owner=root G-Bank

dropdb:
	docker exec -it postgres-1 dropdb G-Bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

dbdocs:
	dbdocs build doc/db.dbml

dbdchema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml
sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/abdulrahman-02/G-Bank/db/sqlc Store

proto:
	rm -f pb/*.go
	rm -f /doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=G-Bank \
	proto/*.proto
	statik -f -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis-1 -p 6379:6379 -d redis:7-alpine

.PHONY: postgres createdb dropdb migrateup  migratedown migrateup1 migratedown1 dbdocs dbdchema \
		sqlc test server mock proto evans redis