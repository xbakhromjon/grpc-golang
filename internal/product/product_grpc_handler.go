package product

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"grpc-golang/protobuf"
	"io"
	"log"
	"time"
)

type ServerImpl struct {
	protobuf.UnimplementedProductServer
}

func (p ServerImpl) Create(ctx context.Context, request *protobuf.ProductRequest) (*protobuf.ProductReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("Error occurred during read metadata")
	}

	log.Println(md.Get("token"))
	log.Println(md.Get("Authorization"))
	log.Println(md.Get("api-key"))
	return &protobuf.ProductReply{
		Name: fmt.Sprintf("Hello %s!", request.Name),
	}, nil
}

func (p ServerImpl) Delete(ctx context.Context, id *protobuf.ProductId) (*protobuf.EmptyReply, error) {
	fmt.Println(id.Id)
	return &protobuf.EmptyReply{}, nil
}

func (p ServerImpl) GetOne(ctx context.Context, id *protobuf.ProductId) (*protobuf.ProductReply, error) {
	fmt.Println(id.Id)
	return &protobuf.ProductReply{
		Name: "John",
	}, nil
}

func (p ServerImpl) Filter(specification *protobuf.ProductFilterSpecification, stream protobuf.Product_FilterServer) error {
	content := []*protobuf.ProductReply{{Name: "John"}, {Name: "Dan"}}
	reply := &protobuf.ProductListReply{
		Content: content,
	}
	for i := 0; i < 5; i++ {
		err := stream.Send(reply)
		if err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}
	return nil
}

func (s ServerImpl) BatchCreate(stream protobuf.Product_BatchCreateServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			err := stream.SendAndClose(&protobuf.EmptyReply{})
			if err != nil {
				return err
			}
			return nil
		}
		log.Printf("%+v", req)
	}
}

func (s ServerImpl) Chat(stream protobuf.Product_ChatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("%+v", req)
		err = stream.Send(&protobuf.ChatContentReply{Content: "content"})
		if err != nil {
			return err
		}
	}

	return nil
}
