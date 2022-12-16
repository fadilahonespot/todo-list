package repository

import (
	"context"

	"github.com/fadilahonespot/todo-list/domain"
	"gorm.io/gorm"
)

type defaultTodo struct {
	db *gorm.DB
}

func SetupRepoActivity(db *gorm.DB) TodoRepository {
	return &defaultTodo{db}
}

func (r *defaultTodo) Create(ctx context.Context, req *domain.Todo) (err error) {
	err = r.db.WithContext(ctx).Create(req).Error
	return
}

func (r *defaultTodo) Update(ctx context.Context, req *domain.Todo) (err error) {
	err = r.db.WithContext(ctx).Save(req).Error
	return
}

func (r *defaultTodo) GetById(ctx context.Context, id int) (resp *domain.Todo, err error) {
	err = r.db.WithContext(ctx).First(&resp, id).Error
	return
}

func (r *defaultTodo) FindAll(ctx context.Context) (resp *[]domain.Todo, err error) {
	err = r.db.WithContext(ctx).Find(&resp).Error
	return
}

func (r *defaultTodo) Delete(ctx context.Context, req *domain.Todo) (err error) {
	err = r.db.WithContext(ctx).Delete(req).Error
	return
}

func (r *defaultTodo) FindByActivityId(ctx context.Context, activityId int) (resp *[]domain.Todo, err error) {
	err = r.db.WithContext(ctx).Where("activity_id = ?", activityId).Find(&resp).Error
	return
}