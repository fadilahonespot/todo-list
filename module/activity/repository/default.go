package repository

import (
	"context"

	"github.com/fadilahonespot/todo-list/domain"
)

type ActivityRepository interface {
	Create(ctx context.Context, req *domain.Activity) (err error)
	Update(ctx context.Context, req *domain.Activity) (err error)
	GetById(ctx context.Context, id int) (resp *domain.Activity, err error)
	FindAll(ctx context.Context) (resp *[]domain.Activity, err error)
	Delete(ctx context.Context, req *domain.Activity) (err error)
}