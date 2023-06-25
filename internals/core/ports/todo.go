package ports

import (
	"context"

	"github.com/jaeyoung0509/todo/internals/core/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, task string) (*domain.Todo, error)
	GetAll(ctx context.Context) (*domain.Todo, error)
	MarkDone(ctx context.Context, id int) (*domain.Todo, error)
	Delete(ctx context.Context, id int) error
}
