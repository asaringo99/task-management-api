package handler

import (
	"encoding/json"
	"net/http"

	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	res "github.com/asaringo99/task_management/http/response"
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
			var response res.ResponseBody
			ret, err := t.controller.Post(c)
			if err != nil {
				response.Status = res.MessageError
				response.Error = err.Error()
				return c.JSON(http.StatusInternalServerError, response)
			}
			response.Status = res.MessageSuccess
			response.Data = ret
			return c.JSON(http.StatusOK, response)
		})

		r.DELETE("/delete/:id", func(c echo.Context) error {
			err := t.controller.Delete(c)
			if err != nil {
				return c.String(http.StatusNoContent, "Error!")
			}
			return c.String(http.StatusOK, "Success")
		})

		r.PUT("/put/:id", func(c echo.Context) error {
			err := t.controller.Put(c)
			if err != nil {
				return c.String(http.StatusNoContent, "Error!")
			}
			return c.String(http.StatusOK, "Success")
		})

		r.PUT("/patch/:id", func(c echo.Context) error {
			err := t.controller.Patch(c)
			if err != nil {
				return c.String(http.StatusNoContent, "Error!")
			}
			return c.String(http.StatusOK, "Success")
		})
	}
}
