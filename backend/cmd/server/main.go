package main

import (
	"net/http"

	"go.uber.org/fx"

	login_controller "github.com/asaringo99/task_management/http/auth/controller/login"
	signup_controller "github.com/asaringo99/task_management/http/auth/controller/signup"
	token_controller "github.com/asaringo99/task_management/http/auth/controller/token"
	auth_model "github.com/asaringo99/task_management/http/auth/model"
	auth_repository "github.com/asaringo99/task_management/http/auth/repository"
	login_usecase "github.com/asaringo99/task_management/http/auth/usecase/login"
	signup_usecase "github.com/asaringo99/task_management/http/auth/usecase/signup"
	token_usecase "github.com/asaringo99/task_management/http/auth/usecase/token"
	"github.com/asaringo99/task_management/http/handler"
	"github.com/asaringo99/task_management/http/router"
	task "github.com/asaringo99/task_management/internal/adapter/controller"
	task_create_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/create"
	task_delete_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/delete"
	task_fetch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/fetch"
	task_create_repository "github.com/asaringo99/task_management/internal/application/repository/task/create"
	task_delete_repository "github.com/asaringo99/task_management/internal/application/repository/task/delete"
	task_fetch_repository "github.com/asaringo99/task_management/internal/application/repository/task/fetch"
	task_create_usecase "github.com/asaringo99/task_management/internal/application/usecase/task/create"
	task_delete_usecase "github.com/asaringo99/task_management/internal/application/usecase/task/delete"
	task_fetch_usecase "github.com/asaringo99/task_management/internal/application/usecase/task/fetch"
	db "github.com/asaringo99/task_management/internal/db"
)

func main() {
	fx.New(
		fx.Provide(
			db.NewTaskManagementConnection,

			token_controller.NewTokenController,
			token_usecase.NewTokenUsecase,
			login_controller.NewLoginController,
			login_usecase.NewLoginUsecase,
			signup_controller.NewSignUpController,
			signup_usecase.NewSignUpUsecase,
			auth_repository.NewAuthRepository,
			auth_model.NewAuthModel,

			task_create_usecase.NewTaskCreateUsecase,
			task_create_usecase.NewTaskCreateInteractor,
			task_create_repository.NewTaskCreateRepository,
			task_create_gateway.NewTaskCreateGateway,

			task_fetch_usecase.NewFetchGetUsecase,
			task_fetch_usecase.NewTaskFetchInteractor,
			task_fetch_repository.NewTaskFetchRepository,
			task_fetch_gateway.NewTaskFetchGateway,

			task_delete_usecase.NewTaskDeleteUsecase,
			task_delete_usecase.NewTaskDeleteInteractor,
			task_delete_repository.NewTaskDeleteRepository,
			task_delete_gateway.NewTaskDeleteGateway,

			task.NewManagementController,

			router.NewEchoServer,
			fx.Annotate(
				router.NewRouter,
				fx.ParamTags(`group:"routes"`),
			),
			handler.AsHandler(handler.NewTaskHandler),
			handler.AsHandler(handler.NewSignUpHandler),
			handler.AsHandler(handler.NewLoginHandler),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
