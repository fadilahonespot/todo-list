package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fadilahonespot/todo-list/module/activity/repository"
	"github.com/fadilahonespot/todo-list/domain"
	"github.com/fadilahonespot/todo-list/utils/cuserr"
)

type defaultActivity struct {
	activityRepo repository.ActivityRepository
}

func SetupUsecaseActivity(activityRepo repository.ActivityRepository) ActivityUsecase {
	return &defaultActivity{activityRepo}
}

func (s *defaultActivity) CreateActivity(ctx context.Context, req *domain.Activity) (err error) {
	err = s.activityRepo.Create(ctx, req)
	if err != nil {
		err = cuserr.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	return
}

func (s *defaultActivity) UpdateActivity(ctx context.Context, id int, req *domain.Activity) (err error) {
	resp, err := s.activityRepo.GetById(ctx, id)
	if err != nil {
		err = cuserr.SetErrorMessage(http.StatusNotFound, fmt.Sprintf("Activity with ID %v Not Found", id))
		return
	}

	req.ID = id
	req.CreatedAt = resp.CreatedAt
	err = s.activityRepo.Update(ctx, req)
	if err != nil {
		err = cuserr.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	return
}

func (s *defaultActivity) FindAllActivity(ctx context.Context) (resp *[]domain.Activity, err error) {
	resp, err = s.activityRepo.FindAll(ctx)
	if err != nil {
		err = cuserr.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	return
}

func (s *defaultActivity) FindById(ctx context.Context, id int) (resp *domain.Activity, err error) {
	resp, err = s.activityRepo.GetById(ctx, id)
	if err != nil {
		err = cuserr.SetErrorMessage(http.StatusNotFound, fmt.Sprintf("Activity with ID %v Not Found", id))
		return
	}

	return
}

func (s *defaultActivity) DeleteActivity(ctx context.Context, id int) (err error) {
	resp, err := s.activityRepo.GetById(ctx, id)
	if err != nil {
		err = cuserr.SetErrorMessage(http.StatusNotFound, fmt.Sprintf("Activity with ID %v Not Found", id))
		return
	}
	err = s.activityRepo.Delete(ctx, resp)
	if err != nil {
		err = cuserr.SetError(http.StatusInternalServerError, err.Error())
	}

	return
}