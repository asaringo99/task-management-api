package router

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func NewEchoServer(lc fx.Lifecycle, e *echo.Echo) (srv *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Start(":8080")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
	return e.Server
}
