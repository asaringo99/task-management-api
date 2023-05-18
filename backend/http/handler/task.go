package handler

import (
	"encoding/json"
	"net/http"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	"github.com/asaringo99/task_management/internal/adapter/controller"
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
		SigningKey: []byte("secret"),
	}
	r.Use(echojwt.WithConfig(config))

	{
		r.GET("/get", func(c echo.Context) error {
			g, err := t.controller.Get(c)
			if err != nil {
				return err
			}
			d, err := json.Marshal(&g)
			if err != nil {
				return c.String(http.StatusNoContent, "Error!")
			}
			return c.JSON(http.StatusOK, echo.Map{
				"response": string(d),
			})
		})

		r.POST("/create", func(c echo.Context) error {
			err := t.controller.Post(c)
			if err != nil {
				return c.String(http.StatusNoContent, "Error!")
			}
			return c.String(http.StatusOK, "Success")
		})

		r.DELETE("/delete/:id", func(c echo.Context) error {
			err := t.controller.Delete(c)
			if err != nil {
				return c.String(http.StatusNoContent, "Error!")
			}
			return c.String(http.StatusOK, "Success")
		})
	}
}
