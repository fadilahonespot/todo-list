package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fadilahonespot/todo-list/domain"
	activityRepo "github.com/fadilahonespot/todo-list/module/activity/repository"
	todoRepo "github.com/fadilahonespot/todo-list/module/todo/repository"
	"github.com/fadilahonespot/todo-list/utils/cuserr"
)

type defaultTodo struct {
	todoRepo     todoRepo.TodoRepository
	activityRepo activityRepo.ActivityRepository
}

func SetupTodoUsecase(todoRepo todoRepo.TodoRepository, activityRepo activityRepo.ActivityRepository) TodoUsecase {
	return &defaultTodo{todoRepo, activityRepo}
}

func (s *defaultTodo) CreateTodo(ctx context.Context, req *domain.Todo) (err error) {
	_, err = s.activityRepo.GetById(ctx, req.ActivityID)
	if err != nil {
		err = cuserr.SetErrorMessage(http.StatusNotFound, fmt.Sprintf("Activity with ID %v Not Found", req.ActivityID))
		return
	}
	err = s.todoRepo.Create(ctx, req)
	if err != nil {
		err = cuserr.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	return
}

func (s *defaultTodo) UpdateTodo(ctx context.Context, req *domain.Todo, id int) (err error) {
	todo, err := s.todoRepo.GetById(ctx, req.ID)
	if err != nil {
		err = cuserr.SetErrorMessage(http.StatusNotFound, fmt.Sprintf("Todo with ID %v Not Found", req.ID))
		return
	}

	_, err = s.activityRepo.GetById(ctx, req.ActivityID)
	if err != nil {
		err = cuserr.SetErrorMessage(http.StatusNotFound, fmt.Sprintf("Activity with ID %v Not Found", req.ActivityID))
		return
	}

	req.ID = id
	req.CreatedAt = todo.CreatedAt
	err = s.todoRepo.Update(ctx, req)
	if err != nil {
		err = cuserr.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	return
}

func (s *defaultTodo) FindAllTodo(ctx context.Context, activityId int) (resp *[]domain.Todo, err error) {
	switch true {
	case activityId != 0:
		activ, errRes := s.todoRepo.FindByActivityId(ctx, activityId)
		if errRes != nil {
			err = cuserr.SetError(http.StatusInternalServerError, errRes.Error())
			return
		}
		resp = activ

	default:
		activ, errRes := s.todoRepo.FindAll(ctx)
		if errRes != nil {
			err = cuserr.SetError(http.StatusInternalServerError, errRes.Error())
			return
		}
		resp = activ
	}

	return
}

func (s *defaultTodo) FindTodoById(ctx context.Context, id int) (resp *domain.Todo, err error) {
	resp, err = s.todoRepo.GetById(ctx, id)
	if err != nil {
		err = cuserr.SetErrorMessage(http.StatusNotFound, fmt.Sprintf("Todo with ID %v Not Found", id))
		return
	}

	return
}

func (s *defaultTodo) DeleteTodo(ctx context.Context, id int) (err error) {
	resp, err := s.todoRepo.GetById(ctx, id)
	if err != nil {
		err = cuserr.SetErrorMessage(http.StatusNotFound, fmt.Sprintf("Todo with ID %v Not Found", id))
		return
	}

	err = s.todoRepo.Delete(ctx, resp)
	if err != nil {
		err = cuserr.SetError(http.StatusInternalServerError, err.Error())
		return
	}

	return
}
