package server

import (
	"context"
	"fmt"
	"github.com/kkmmttdd/todo-grpc/src/delivery/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {}

func (*server) TodoList(ctx context.Context, req *todo_grpc.TodoListRequest) (*todo_grpc.TodoListResponse, error) {
	dummy_todo := todo_grpc.Todo{Text: "hogehoge"}
	return &todo_grpc.TodoListResponse{
		Todos: []*todo_grpc.Todo{&dummy_todo},
	}, nil
}

func Run() {
	fmt.Println("Server Start")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	todo_grpc.RegisterTodoServiceServer(s, &server{})


	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("%v", err)
	}
}

