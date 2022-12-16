package main

import (
	"log"

	"github.com/fadilahonespot/todo-list/utils/cuserr"
	"github.com/fadilahonespot/todo-list/utils/database"
	"github.com/fadilahonespot/todo-list/router"
	activityRepo "github.com/fadilahonespot/todo-list/module/activity/repository"
	activityUsecase "github.com/fadilahonespot/todo-list/module/activity/usecase"
	activityHandler "github.com/fadilahonespot/todo-list/module/activity/handler"
	todoRepo "github.com/fadilahonespot/todo-list/module/todo/repository"
	todoUsecase "github.com/fadilahonespot/todo-list/module/todo/usecase"
	todoHandler "github.com/fadilahonespot/todo-list/module/todo/handler"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Setup database
	db := database.NewDatabase()

	// Setup Module
	actRepo := activityRepo.SetupRepoActivity(db)
	tdRepo := todoRepo.SetupRepoActivity(db)

	actUsecase := activityUsecase.SetupUsecaseActivity(actRepo)
	tdUsecase := todoUsecase.SetupTodoUsecase(tdRepo, actRepo)

	actHandler := activityHandler.SetupActivityHandler(actUsecase)
	tdHandler := todoHandler.SetupActivityHandler(tdUsecase)

	// Setup Router
	app := fiber.New(
		fiber.Config{
			ErrorHandler: cuserr.ErrorHandle,
		},
	)

	router.SetupRouter(*actHandler, *tdHandler).SetRouter(app)
	log.Fatal(app.Listen(":3030"))
}