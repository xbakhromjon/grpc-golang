package product

import (
	"context"
	"fmt"
)

type ServerImpl struct {
	UnimplementedProductServer
}

func (p ServerImpl) Create(ctx context.Context, request *ProductRequest) (*ProductReply, error) {

	return &ProductReply{
		Name: fmt.Sprintf("Hello %s!", request.Name),
	}, nil
}

func (p ServerImpl) Delete(ctx context.Context, id *ProductId) (*EmptyReply, error) {
	fmt.Println(id.Id)
	return &EmptyReply{}, nil
}

func (p ServerImpl) GetOne(ctx context.Context, id *ProductId) (*ProductReply, error) {
	fmt.Println(id.Id)
	return &ProductReply{
		Name: "John",
	}, nil
}

func (p ServerImpl) Filter(ctx context.Context, specification *ProductFilterSpecification) (*ProductListReply, error) {

	content := []*ProductReply{{Name: "John"}, {Name: "Dan"}}
	return &ProductListReply{
		Content: content,
	}, nil
}
