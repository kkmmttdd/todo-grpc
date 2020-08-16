package todo

import (
	"context"
	deliveryGrpc "github.com/kkmmttdd/todo-grpc/src/delivery/grpc"
	todoUsecase "github.com/kkmmttdd/todo-grpc/src/usecase/todo"
	todoAdapter "github.com/kkmmttdd/todo-grpc/src/util/grpc/adapter/models/todo"
)

type TodoServer struct {

}

func (t *TodoServer) TodoList(ctx context.Context, req *deliveryGrpc.TodoListRequest) (*deliveryGrpc.TodoListResponse, error) {
	todos := todoUsecase.UC.Search()
	var todosRes []*deliveryGrpc.Todo
	for _, td := range *todos {
		td_res := todoAdapter.DomainToGrpc{td}.Adapt()
		todosRes = append(todosRes, td_res)
	}
	res := deliveryGrpc.TodoListResponse{Todos:todosRes}
	return &res, nil
}

func (t *TodoServer) TodoCreate(ctx context.Context, req *deliveryGrpc.TodoCreateRequest) (*deliveryGrpc.TodoCreateResponse, error) {
	requestDomain := todoAdapter.GrpcToDomain{req.Todo}.Adapt()
	valid, _ := todoUsecase.UC.Create(requestDomain)
	return &deliveryGrpc.TodoCreateResponse{Valid:valid, Message: ""}, nil
}
