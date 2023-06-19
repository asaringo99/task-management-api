package handler

import (
	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	controller "github.com/asaringo99/task_management/internal/adapter/controller/task"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	controller controller.TaskController
}

func NewTaskHandler(controller *controller.TaskController) *TaskHandler {
	return &TaskHandler{
		controller: *controller,
	}
}

func (t *TaskHandler) AddEntryPoint(e *echo.Echo) {
	r := e.Group("/task")

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(authjwt.JwtCustomClaims)
		},
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:token",
	}
	r.Use(echojwt.WithConfig(config))

	{
		r.GET("/get", t.get)

		r.POST("/create", t.post)

		r.DELETE("/delete/:id", t.delete)

		r.PUT("/put/:id", t.put)

		r.PUT("/patch/:id", t.patch)
	}
}
