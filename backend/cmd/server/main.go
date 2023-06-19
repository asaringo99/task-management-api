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
	board_handler "github.com/asaringo99/task_management/http/handler/board"
	task_handler "github.com/asaringo99/task_management/http/handler/task"
	"github.com/asaringo99/task_management/http/router"
	board_controller "github.com/asaringo99/task_management/internal/adapter/controller/board"
	task_controller "github.com/asaringo99/task_management/internal/adapter/controller/task"
	board_create_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/create"
	board_delete_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/delete"
	board_fetch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/fetch"
	board_put_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/put"
	task_create_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/create"
	task_delete_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/delete"
	task_fetch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/fetch"
	task_patch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/patch"
	task_put_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/put"
	board_create_repository "github.com/asaringo99/task_management/internal/application/repository/board/create"
	board_delete_repository "github.com/asaringo99/task_management/internal/application/repository/board/delete"
	board_fetch_repository "github.com/asaringo99/task_management/internal/application/repository/board/fetch"
	board_put_repository "github.com/asaringo99/task_management/internal/application/repository/board/put"
	task_create_repository "github.com/asaringo99/task_management/internal/application/repository/task/create"
	task_delete_repository "github.com/asaringo99/task_management/internal/application/repository/task/delete"
	task_fetch_repository "github.com/asaringo99/task_management/internal/application/repository/task/fetch"
	task_patch_repository "github.com/asaringo99/task_management/internal/application/repository/task/patch"
	task_put_repository "github.com/asaringo99/task_management/internal/application/repository/task/put"
	board_create_usecase "github.com/asaringo99/task_management/internal/application/usecase/board/create"
	board_delete_usecase "github.com/asaringo99/task_management/internal/application/usecase/board/delete"
	board_fetch_usecase "github.com/asaringo99/task_management/internal/application/usecase/board/fetch"
	board_put_usecase "github.com/asaringo99/task_management/internal/application/usecase/board/put"
	task_create_usecase "github.com/asaringo99/task_management/internal/application/usecase/task/create"
	task_delete_usecase "github.com/asaringo99/task_management/internal/application/usecase/task/delete"
	task_fetch_usecase "github.com/asaringo99/task_management/internal/application/usecase/task/fetch"
	task_patch_usecase "github.com/asaringo99/task_management/internal/application/usecase/task/patch"
	task_put_usecase "github.com/asaringo99/task_management/internal/application/usecase/task/put"
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

			task_put_usecase.NewTaskPutUsecase,
			task_put_usecase.NewTaskPutInteractor,
			task_put_repository.NewTaskPutRepository,
			task_put_gateway.NewTaskPutGateway,

			task_patch_usecase.NewTaskPatchUsecase,
			task_patch_usecase.NewTaskPatchInteractor,
			task_patch_repository.NewTaskPatchRepository,
			task_patch_gateway.NewTaskPatchGateway,

			task_controller.NewManagementController,

			board_create_usecase.NewBoardCreateUsecase,
			board_create_usecase.NewBoardCreateInteractor,
			board_create_repository.NewBoardCreateRepository,
			board_create_gateway.NewBoardCreateGateway,

			board_fetch_usecase.NewFetchGetUsecase,
			board_fetch_usecase.NewBoardFetchInteractor,
			board_fetch_repository.NewBoardFetchRepository,
			board_fetch_gateway.NewBoardFetchGateway,

			board_delete_usecase.NewBoardDeleteUsecase,
			board_delete_usecase.NewBoardDeleteInteractor,
			board_delete_repository.NewBoardDeleteRepository,
			board_delete_gateway.NewBoardDeleteGateway,

			board_put_usecase.NewBoardPutUsecase,
			board_put_usecase.NewBoardPutInteractor,
			board_put_repository.NewBoardPutRepository,
			board_put_gateway.NewBoardPutGateway,

			board_controller.NewBoardController,

			router.NewEchoServer,
			fx.Annotate(
				router.NewRouter,
				fx.ParamTags(`group:"routes"`),
			),
			handler.AsHandler(task_handler.NewTaskHandler),
			handler.AsHandler(board_handler.NewBoardHandler),
			handler.AsHandler(handler.NewSignUpHandler),
			handler.AsHandler(handler.NewLoginHandler),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
