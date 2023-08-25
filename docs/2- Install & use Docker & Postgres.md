# Install & use Docker & Postgres

## What we will do

- Install Docker Desktop
- Run Postgres Container
- Install Table Plus
- Create the DB schema

## Docker Commands

- `docker pull postgres`
- `docker images`
- `docker ps`
- `docker run --name postgres-1 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
- `docker exec -it postgres-1 psql -U root`
- `docker logs postgres-`
