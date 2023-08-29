# How to connect containers in the same docker network

- `docker run --name g-bank -p 8080:8080 -e GIN_MODE=release g-bank:latest`
- `docker run --name g-bank-db -p 5432:5432 -eGIN_MODE=release -e DB_SOURCE="postgres://root:password@172.17.0.3:5432/G-Bank?sslmode=disable" g-bank:latest`
- `docker network ls`
- `docker network inspect bridge`
- `docker network create g-bank-network`
- `docker network connect g-bank-network postgres-1`
- `docker run --name g-bank-db --network g-bank-network -p 5432:5432 -eGIN_MODE=release -e DB_SOURCE="postgres://root:password@1postgres-1:5432/G-Bank?sslmode=disable" g-bank:latest`
