package handler

import (
	"net/http"

	"github.com/fadilahonespot/todo-list/module/activity/usecase"
	"github.com/fadilahonespot/todo-list/utils/response"
	"github.com/fadilahonespot/todo-list/domain"
	"github.com/fadilahonespot/todo-list/utils/cuserr"
	"github.com/gofiber/fiber/v2"
)

type ActivityHandler struct {
	activityUsecase usecase.ActivityUsecase
}

func SetupActivityHandler(activityUsecase usecase.ActivityUsecase) *ActivityHandler {
	return &ActivityHandler{activityUsecase}
}

func (h *ActivityHandler) CreateHandler(c *fiber.Ctx) (err error) {
	var req domain.Activity
	err = c.BodyParser(&req)
	if err != nil {
		err = cuserr.SetError(http.StatusBadRequest, err.Error())
		return
	}
	if req.Title == "" {
		err = cuserr.SetErrorMessage(http.StatusBadRequest, "title cannot be null")
		return
	}
	err = h.activityUsecase.CreateActivity(c.Context(), &req)
	if err != nil {
		return
	}

	return response.HandleSuccess(c, req)
}

func (h *ActivityHandler) UpdateActivity(c *fiber.Ctx) (err error) {
	var req domain.Activity
	err = c.BodyParser(&req)
	if err != nil {
		err = cuserr.SetError(http.StatusBadRequest, err.Error())
		return
	}
	id, _ := c.ParamsInt("id")
	if req.Title == "" {
		err = cuserr.SetErrorMessage(http.StatusBadRequest, "title cannot be null")
		return
	}
	err = h.activityUsecase.UpdateActivity(c.Context(), id, &req)
	if err != nil {
		return
	}

	return response.HandleSuccess(c, req)
}

func (h *ActivityHandler) FindAllActivity(c *fiber.Ctx) (err error) {
	resp, err := h.activityUsecase.FindAllActivity(c.Context())
	if err != nil {
		return
	}

	return response.HandleSuccess(c, resp)
}

func (h *ActivityHandler) FindActivityById(c *fiber.Ctx) (err error) {
	id, _ := c.ParamsInt("id")
	
	resp, err := h.activityUsecase.FindById(c.Context(), id)
	if err != nil {
		return
	}

	return response.HandleSuccess(c, resp)
}

func (h *ActivityHandler) DeleteActivity(c *fiber.Ctx) (err error) {
	id, _ := c.ParamsInt("id")
	err = h.activityUsecase.DeleteActivity(c.Context(), id)
	if err != nil {
		return
	}

	return response.HandleSuccess(c, struct{}{})
}