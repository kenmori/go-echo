package main

import (
	"go-echo/controller"
	"go-echo/db"
	"go-echo/repository"
	"go-echo/router"
	"go-echo/usecase"
	"go-echo/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUseCase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUseCase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUseCase)
	taskController := controller.NewTaskController(taskUseCase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
