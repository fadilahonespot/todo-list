package usecase

import (
	"context"

	"github.com/fadilahonespot/todo-list/domain"
)

type ActivityUsecase interface {
	CreateActivity(ctx context.Context, req *domain.Activity) (err error)
	UpdateActivity(ctx context.Context, id int, req *domain.Activity) (err error)
	FindAllActivity(ctx context.Context) (resp *[]domain.Activity, err error)
	FindById(ctx context.Context, id int) (resp *domain.Activity, err error)
	DeleteActivity(ctx context.Context, id int) (err error)
}

