package todo

import (
	"context"
	"time"
)

type Todo struct {
	Id          string    `json:"id"`
	DateCreated time.Time `json:"dateCreated"`
	Text        string    `json:"text"`
	Completed   bool      `json:"completed"`
}

func (o *Service) ListTodos(ctx context.Context) ([]Todo, error) {
	return []Todo{
		{"123", time.Now(), "this is totally a real todo and this is what it looks like when it gets longer", false},
	}, nil
}
