package repository

import (
	"context"

	"github.com/fadilahonespot/todo-list/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, req *domain.Todo) (err error)
	Update(ctx context.Context, req *domain.Todo) (err error)
	GetById(ctx context.Context, id int) (resp *domain.Todo, err error)
	FindAll(ctx context.Context) (resp *[]domain.Todo, err error)
	Delete(ctx context.Context, req *domain.Todo) (err error)
	FindByActivityId(ctx context.Context, activityId int) (resp *[]domain.Todo, err error)
}