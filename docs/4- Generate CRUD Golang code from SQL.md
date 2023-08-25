# Generate CRUD Golang code form SQL

## What is CRUD?

- CREATE : insert to the db
- READ : select from the db
- UPDATE : change some fields in the db
- DELETE :  remove from the db

## Things to consider

- DATABASE/SQL :  <https://pkg.go.dev/database/sql>
  - very fast & straightforward
  - Manual mapping SQL fields to variables
  - Easy to make mistakes, not caught until runtime
- GORM : <https://gorm.io/index.html>
  - CRUD functions already implemented, very short production code
  - must learn how to write queries using gorm's functions
  - Run slowly on high load
- SQLX : <https://github.com/jmoiron/sqlx>
  - Quite fast & easy to use
  - Fields mapping via query text & struct tags
  - Failure won't occur until runtime
- SQLC :
  - very fast & easy to use
  - automatic code generation
  - Full support Postgres.

## SQLC

- `sqlc init`
- `sqlc generate`
