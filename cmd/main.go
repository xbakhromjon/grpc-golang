package main

import (
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-golang/internal/auth"
	"grpc-golang/internal/product"
	"grpc-golang/protobuf"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	go func() {
		err := http.ListenAndServe(":5006", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("ping \n"))
		}))

		if err != nil {
			log.Println(err)
		}
	}()

	// Load server's TLS certificate and private key
	//serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Create a TLS configuration for the server
	//tlsConfig := &tls.Config{
	//	Certificates: []tls.Certificate{serverCert},
	//}

	lis, err := net.Listen("tcp", os.Getenv("RPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server with TLS
	//s := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))
	s := grpc.NewServer()
	protobuf.RegisterProductServer(s, &product.ServerImpl{})
	protobuf.RegisterAuthServer(s, &auth.ServerImpl{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Server started on :5005")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
