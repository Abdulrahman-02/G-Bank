# Codespaces

```bash
# Pull Postgres
docker pull postgres
# Install migrate
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 /usr/local/bin/migrate
# Install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# Install mockgen
go install go.uber.org/mock/mockgen@latest
# Install dbdocs
npm install -g dbdocs
# Install dbml cli
npm install -g @dbml/cli
# Install protobuf compiler
apt install -y protobuf-compiler
# Install Go protobuf compiler plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# Install Go gRPC plugin
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# Add $GOPATH/bin to your PATH
export PATH="$PATH:$(go env GOPATH)/bin"

```
