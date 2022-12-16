package router

import (
	activityHandler "github.com/fadilahonespot/todo-list/module/activity/handler"
	todoHandler "github.com/fadilahonespot/todo-list/module/todo/handler"
	"github.com/gofiber/fiber/v2"
)

type defaultRouter struct {
	actHandler activityHandler.ActivityHandler
	tdHandler  todoHandler.TodoHandler
}

func SetupRouter(activityHandler activityHandler.ActivityHandler, tdHandler todoHandler.TodoHandler) *defaultRouter {
	return &defaultRouter{activityHandler, tdHandler}
}

func (r *defaultRouter) SetRouter(app *fiber.App) {
	app.Get("/activity-groups", r.actHandler.FindAllActivity)
	app.Get("/activity-groups/:id", r.actHandler.FindActivityById)
	app.Post("/activity-groups", r.actHandler.CreateHandler)
	app.Delete("/activity-groups/:id", r.actHandler.DeleteActivity)
	app.Patch("/activity-groups/:id", r.actHandler.UpdateActivity)

	app.Get("/todo-items", r.tdHandler.FindAllTodo)
	app.Get("/todo-items/:id", r.tdHandler.FindTodoById)
	app.Post("/todo-items", r.tdHandler.CreateTodo)
	app.Delete("/todo-items/:id", r.tdHandler.DeleteTodo)
	app.Patch("/todo-items/:id", r.tdHandler.UpdateTodo)
}
