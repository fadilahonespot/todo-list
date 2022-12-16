package repository

import (
	"context"

	"github.com/fadilahonespot/todo-list/domain"
	"gorm.io/gorm"
)

type defaultActivity struct {
	db *gorm.DB
}

func SetupRepoActivity(db *gorm.DB) ActivityRepository {
	return &defaultActivity{db}
}

func (r *defaultActivity) Create(ctx context.Context, req *domain.Activity) (err error) {
	err = r.db.WithContext(ctx).Create(req).Error
	return
}

func (r *defaultActivity) Update(ctx context.Context, req *domain.Activity) (err error) {
	err = r.db.WithContext(ctx).Save(req).Error
	return
}

func (r *defaultActivity) GetById(ctx context.Context, id int) (resp *domain.Activity, err error) {
	err = r.db.WithContext(ctx).First(&resp, id).Error
	return
}

func (r *defaultActivity) FindAll(ctx context.Context) (resp *[]domain.Activity, err error) {
	err = r.db.WithContext(ctx).Find(&resp).Error
	return
}

func (r *defaultActivity) Delete(ctx context.Context, req *domain.Activity) (err error) {
	err = r.db.WithContext(ctx).Delete(req).Error
	return
}