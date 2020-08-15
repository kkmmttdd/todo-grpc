package todo

import (
	"github.com/kkmmttdd/todo-grpc/src/domain"
	"github.com/kkmmttdd/todo-grpc/src/repository/todo"
)

var UC UseCase

func init()  {
	UC = UseCase{}
}
type UseCase struct {
}

func (u UseCase) Search() *[]*domain.Todo {
	return todo.Repo.Search()
}


