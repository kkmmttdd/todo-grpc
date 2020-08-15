package server

import (
	"fmt"
	deliveryGrpc "github.com/kkmmttdd/todo-grpc/src/delivery/grpc"
	"github.com/kkmmttdd/todo-grpc/src/delivery/grpc/server/servers/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func Run() {
	fmt.Println("Server Start")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	deliveryGrpc.RegisterTodoServiceServer(s, &todo.TodoServer{})


	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("%v", err)
	}
}

