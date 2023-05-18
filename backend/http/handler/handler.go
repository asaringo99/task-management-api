package handler

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func AsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"routes"`),
	)
}

type Handler interface {
	AddEntryPoint(*echo.Echo)
}
