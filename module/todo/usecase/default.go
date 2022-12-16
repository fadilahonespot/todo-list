package usecase

import (
	"context"

	"github.com/fadilahonespot/todo-list/domain"
)

type TodoUsecase interface {
	CreateTodo(ctx context.Context, req *domain.Todo) (err error)
	UpdateTodo(ctx context.Context, req *domain.Todo, id int) (err error)
	FindAllTodo(ctx context.Context, activityId int) (resp *[]domain.Todo, err error)
	FindTodoById(ctx context.Context, id int) (resp *domain.Todo, err error)
	DeleteTodo(ctx context.Context, id int) (err error) 
}