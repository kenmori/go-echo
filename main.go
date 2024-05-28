package main

import (
	"go-echo/controller"
	"go-echo/db"
	"go-echo/repository"
	"go-echo/router"
	"go-echo/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUseCase := usecase.NewUserUsecase(userRepository)
	taskUseCase := usecase.NewTaskUsecase(taskRepository)

	userController := controller.NewUserController(userUseCase)
	taskController := controller.NewTaskController(taskUseCase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
