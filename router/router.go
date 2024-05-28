package router

import (
	"go-echo/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.LogOut)
	t := e.Group("/tasks")
	// jwtのmiddlewareを使って認証を行う。USEを使うことで、tasksグループ内の全てのエンドポイントに認証をかけることができる
	t.Use(echojwt.WithConfig(echojwt.Config{
		// dbを生成した時と同じSECRETを使う
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
