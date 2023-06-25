package service

import (
	"context"

	"github.com/jaeyoung0509/todo/internals/core/domain"
	"github.com/jaeyoung0509/todo/internals/core/ports"
)

type TodoService struct {
	Repo ports.TodoRepository
}

func NewTodoService(repo ports.TodoRepository) *TodoService {
	return &TodoService{
		Repo: repo,
	}
}

func (s *TodoService) Create(ctx context.Context, t string) (*domain.Todo, error) {
	return s.Repo.Create(ctx, t)
}

func (s *TodoService) GetAll(ctx context.Context, t string) (*domain.Todo, error) {
	return s.Repo.GetAll(ctx)
}

func (s *TodoService) MarkDone(ctx context.Context, id int) (*domain.Todo, error) {
	return s.Repo.MarkDone(ctx, id)
}
func (s *TodoService) Delete(ctx context.Context, id int) error {
	return s.Repo.Delete(ctx, id)
}
