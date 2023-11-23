1. **Install the `protoc` Compiler:** https://grpc.io/docs/protoc-installation/
2. **Install the `protoc-gen-go` Plugin:** `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
3. `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
4. `go get -u google.golang.org/grpc`
5. **Add `protoc-gen-go` to Your PATH: `export PATH=$PATH:$(go env GOPATH)/bin`**
6. `protoc --go_out=. --go-grpc_out=. your_proto_file.proto`
7. xxx.pb.go va xxx_grpc_pb.go fayllarni generate qilib beradi.