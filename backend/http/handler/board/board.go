package handler

import (
	authjwt "github.com/asaringo99/task_management/http/auth/jwt"
	controller "github.com/asaringo99/task_management/internal/adapter/controller/board"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type BoardHandler struct {
	controller controller.BoardController
}

func NewBoardHandler(controller *controller.BoardController) *BoardHandler {
	return &BoardHandler{
		controller: *controller,
	}
}

func (t *BoardHandler) AddEntryPoint(e *echo.Echo) {
	r := e.Group("/board")

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

		r.PATCH("/patch/:id", t.patch)
	}
}
