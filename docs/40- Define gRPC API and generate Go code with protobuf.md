# Define gRPC API and generate Go code with protobuf

<https://grpc.io/docs>

## Quick Start

```bash
# Install protobuf compiler
apt install -y protobuf-compiler
# Install Go protobuf compiler plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# Install Go gRPC plugin
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
# Add $GOPATH/bin to your PATH
export PATH="$PATH:$(go env GOPATH)/bin"
```
