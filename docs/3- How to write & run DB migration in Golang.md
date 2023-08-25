# How to write & run DB migration in Golang

- Install  <https://github.com/golang-migrate/migrate>
- `migrate create -ext sql -dir db/migration -seq init_schema`
- Generated 2 files up & down
- Created Makefile to automate docker commands
- Start the migration `migrate -path db/migration -database "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable" -verbose up`

## Why ?

![migrateUp](img/db-migration.png)

## Migrate Up

![migrateUp](img/migrateUp.png)

## Migrate Down

![migrateDown](img/migrateDown.png)
