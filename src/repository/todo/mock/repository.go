package mock

import "github.com/kkmmttdd/todo-grpc/src/domain"

type Repository struct {

}

func (r Repository) Search() *[]*domain.Todo {
	return &[]*domain.Todo{
		{Text: "Wasshoi!"},
	}
}

