package mock

import "github.com/kkmmttdd/todo-grpc/src/domain"

type Repository struct {

}

func (r Repository) Search(userId int16) *[]*domain.Todo {
	all := []*domain.Todo{
		{
			Text: "111111111",
			UserId: 1,
		},
		{
			Text: "222222222",
			UserId: 2,
		},
	}
	filtered := []*domain.Todo{}
	for _, td := range all {
		if td.UserId == userId {
			filtered = append(filtered, td)
		}
	}
	return &filtered
}

func (r Repository) Create(t *domain.Todo) (bool, error) {
	return true, nil
}
