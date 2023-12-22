package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-golang/protobuf"
	"log"
)

type ServerImpl struct {
	protobuf.UnimplementedAuthServer
}

func (s *ServerImpl) Signup(context context.Context, request *protobuf.SignupRequest) (*protobuf.EmptyReply, error) {
	log.Printf("request body: %+v", request)
	err := grpc.SetHeader(context, metadata.New(map[string]string{"custom-header": "header-val"}))
	if err != nil {
		return nil, err
	}
	err = grpc.SetTrailer(context, metadata.New(map[string]string{"custom-trailer": "trailer-val"}))
	if err != nil {
		return nil, err
	}
	return nil, status.Errorf(codes.InvalidArgument, "Invalid argument provided")
	return &protobuf.EmptyReply{}, nil
}

func (s *ServerImpl) SignIn(context context.Context, request *protobuf.SignInRequest) (*protobuf.TokenReply, error) {
	log.Printf("request body: %+v", request)
	return &protobuf.TokenReply{Token: "token123"}, nil
}
