package handler

import (
	"net/http"
	"strconv"

	"github.com/fadilahonespot/todo-list/domain"
	"github.com/fadilahonespot/todo-list/module/todo/usecase"
	"github.com/fadilahonespot/todo-list/utils/cuserr"
	"github.com/fadilahonespot/todo-list/utils/response"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func SetupActivityHandler(todoUsecase usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{todoUsecase}
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) (err error) {
	var req domain.Todo
	err = c.BodyParser(&req)
	if err != nil {
		err = cuserr.SetError(http.StatusBadRequest, err.Error())
		return
	}
	if req.Title == "" {
		err = cuserr.SetErrorMessage(http.StatusBadRequest, "title cannot be null")
		return
	}
	err = h.todoUsecase.CreateTodo(c.Context(), &req)
	if err != nil {
		return
	}

	return response.HandleSuccess(c, req)
}

func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) (err error) {
	var req domain.Todo
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
	err = h.todoUsecase.UpdateTodo(c.Context(), &req, id)
	if err != nil {
		return
	}

	return response.HandleSuccess(c, req)
}

func (h *TodoHandler) FindAllTodo(c *fiber.Ctx) (err error) {
	activityId := c.Query("activity_group_id")
	id, _ := strconv.Atoi(activityId)
	resp, err := h.todoUsecase.FindAllTodo(c.Context(), id)
	if err != nil {
		return
	}

	return response.HandleSuccess(c, resp)
}

func (h *TodoHandler) FindTodoById(c *fiber.Ctx) (err error) {
	id, _ := c.ParamsInt("id")
	
	resp, err := h.todoUsecase.FindTodoById(c.Context(), id)
	if err != nil {
		return
	}

	return response.HandleSuccess(c, resp)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) (err error) {
	id, _ := c.ParamsInt("id")
	err = h.todoUsecase.DeleteTodo(c.Context(), id)
	if err != nil {
		return
	}

	return response.HandleSuccess(c, struct{}{})
}