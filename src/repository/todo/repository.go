package todo

import (
	"github.com/kkmmttdd/todo-grpc/src/domain"
	"github.com/kkmmttdd/todo-grpc/src/repository/todo/mock"
)

var (
	Repo Repository
)

func init() {
	Repo = mock.Repository{}
}

type Repository interface {
	Search() *[]*domain.Todo
	Create(*domain.Todo) (bool, error)
}

