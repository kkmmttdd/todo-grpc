package todo

import (
	"github.com/kkmmttdd/todo-grpc/src/delivery/grpc"
	"github.com/kkmmttdd/todo-grpc/src/domain"
)

type GrpcToDomain struct {
	GrpcModel *grpc.Todo
}

func (G GrpcToDomain) Adapt() *domain.Todo {
	return &domain.Todo{Text:"waaaaaaaaaaaay"}
}

type DomainToGrpc struct {
	DomainModel *domain.Todo
}

func (G DomainToGrpc) Adapt() *grpc.Todo {
	todo := grpc.Todo{}
	todo.Text = G.DomainModel.Text
	return &todo
}
