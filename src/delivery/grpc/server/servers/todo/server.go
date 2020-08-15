package todo

import (
	"context"
	deliveryGrpc "github.com/kkmmttdd/todo-grpc/src/delivery/grpc"
	todoUsecase "github.com/kkmmttdd/todo-grpc/src/usecase/todo"
	"github.com/kkmmttdd/todo-grpc/src/util/grpc/adapter/models/todo"
)

type TodoServer struct {

}

func (t *TodoServer) TodoList(ctx context.Context, req *deliveryGrpc.TodoListRequest) (*deliveryGrpc.TodoListResponse, error) {
	todos := todoUsecase.UC.Search()
	var todosRes []*deliveryGrpc.Todo
	for _, td := range *todos {
		td_res := todo.DomainToGrpc{td}.Adapt()
		todosRes = append(todosRes, td_res)
	}
	res := deliveryGrpc.TodoListResponse{Todos:todosRes}
	return &res, nil
}
