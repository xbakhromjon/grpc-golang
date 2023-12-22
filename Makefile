cert:
	cd cert; ./gen-server-cert.sh; ./gen-client-cert.sh; cd ..;

start:
	go run cmd/main.go

proto-gen:
	 protoc --go_out=. --go-grpc_out=. proto/product.proto proto/common.proto proto/auth.proto
