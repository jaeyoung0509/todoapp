package repositories

import (
	"context"

	"github.com/jaeyoung0509/todo/ent"
	"github.com/jaeyoung0509/todo/internals/core/domain"
	"github.com/jaeyoung0509/todo/internals/core/ports"
)

type TodoAdapter struct {
	Client *ent.Client
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.TodoRepository = (*TodoAdapter)(nil)

func NewTodoAdapter(client *ent.Client) *TodoAdapter {
	return &TodoAdapter{
		Client: client,
	}
}

func (adapter *TodoAdapter) Create(ctx context.Context, t string) (*domain.Todo, error) {
	_, err := adapter.Client.Todo.Create().SetTask(t).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.Todo{}, nil
}

func (adapter *TodoAdapter) GetAll(ctx context.Context) (*domain.Todo, error) {
	_, err := adapter.Client.Todo.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.Todo{}, nil
}

func (adapter *TodoAdapter) MarkDone(ctx context.Context, id int) (*domain.Todo, error) {
	_, err := adapter.Client.Todo.
		UpdateOneID(id).
		SetDone(true).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.Todo{}, nil
}

func (adapter *TodoAdapter) Delete(ctx context.Context, id int) error {
	_ = adapter.Client.Todo.
		DeleteOneID(id).
		Exec(ctx)

	return nil
}
