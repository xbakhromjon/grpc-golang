package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-golang/product"
	"log"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	client := product.NewProductClient(conn)

	message := product.ProductFilterSpecification{Name: "John"}
	response, err := client.Filter(context.Background(), &message)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("response: %+v", response)
}
