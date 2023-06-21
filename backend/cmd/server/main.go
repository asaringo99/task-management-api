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
	tab_handler "github.com/asaringo99/task_management/http/handler/tab"
	task_handler "github.com/asaringo99/task_management/http/handler/task"
	"github.com/asaringo99/task_management/http/router"
	board_controller "github.com/asaringo99/task_management/internal/adapter/controller/board"
	tab_controller "github.com/asaringo99/task_management/internal/adapter/controller/tab"
	task_controller "github.com/asaringo99/task_management/internal/adapter/controller/task"
	board_create_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/create"
	board_delete_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/delete"
	board_fetch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/fetch"
	board_patch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/patch"
	board_put_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/board/put"
	tab_create_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/create"
	tab_delete_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/delete"
	tab_fetch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/fetch"
	tab_patch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/patch"
	tab_put_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/tab/put"
	task_create_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/create"
	task_delete_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/delete"
	task_fetch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/fetch"
	task_patch_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/patch"
	task_put_gateway "github.com/asaringo99/task_management/internal/adapter/gateway/task/put"
	board_create_repository "github.com/asaringo99/task_management/internal/application/repository/board/create"
	board_delete_repository "github.com/asaringo99/task_management/internal/application/repository/board/delete"
	board_fetch_repository "github.com/asaringo99/task_management/internal/application/repository/board/fetch"
	board_patch_repository "github.com/asaringo99/task_management/internal/application/repository/board/patch"
	board_put_repository "github.com/asaringo99/task_management/internal/application/repository/board/put"
	tab_create_repository "github.com/asaringo99/task_management/internal/application/repository/tab/create"
	tab_delete_repository "github.com/asaringo99/task_management/internal/application/repository/tab/delete"
	tab_fetch_repository "github.com/asaringo99/task_management/internal/application/repository/tab/fetch"
	tab_patch_repository "github.com/asaringo99/task_management/internal/application/repository/tab/patch"
	tab_put_repository "github.com/asaringo99/task_management/internal/application/repository/tab/put"
	task_create_repository "github.com/asaringo99/task_management/internal/application/repository/task/create"
	task_delete_repository "github.com/asaringo99/task_management/internal/application/repository/task/delete"
	task_fetch_repository "github.com/asaringo99/task_management/internal/application/repository/task/fetch"
	task_patch_repository "github.com/asaringo99/task_management/internal/application/repository/task/patch"
	task_put_repository "github.com/asaringo99/task_management/internal/application/repository/task/put"
	board_create_usecase "github.com/asaringo99/task_management/internal/application/usecase/board/create"
	board_delete_usecase "github.com/asaringo99/task_management/internal/application/usecase/board/delete"
	board_fetch_usecase "github.com/asaringo99/task_management/internal/application/usecase/board/fetch"
	board_patch_usecase "github.com/asaringo99/task_management/internal/application/usecase/board/patch"
	board_put_usecase "github.com/asaringo99/task_management/internal/application/usecase/board/put"
	tab_create_usecase "github.com/asaringo99/task_management/internal/application/usecase/tab/create"
	tab_delete_usecase "github.com/asaringo99/task_management/internal/application/usecase/tab/delete"
	tab_fetch_usecase "github.com/asaringo99/task_management/internal/application/usecase/tab/fetch"
	tab_patch_usecase "github.com/asaringo99/task_management/internal/application/usecase/tab/patch"
	tab_put_usecase "github.com/asaringo99/task_management/internal/application/usecase/tab/put"
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

			board_patch_usecase.NewBoardPatchUsecase,
			board_patch_usecase.NewBoardPatchInteractor,
			board_patch_repository.NewBoardPatchRepository,
			board_patch_gateway.NewBoardPatchGateway,

			board_controller.NewBoardController,

			tab_create_usecase.NewTabCreateUsecase,
			tab_create_usecase.NewTabCreateInteractor,
			tab_create_repository.NewTabCreateRepository,
			tab_create_gateway.NewTabCreateGateway,

			tab_fetch_usecase.NewFetchGetUsecase,
			tab_fetch_usecase.NewTabFetchInteractor,
			tab_fetch_repository.NewTabFetchRepository,
			tab_fetch_gateway.NewTabFetchGateway,

			tab_delete_usecase.NewTabDeleteUsecase,
			tab_delete_usecase.NewTabDeleteInteractor,
			tab_delete_repository.NewTabDeleteRepository,
			tab_delete_gateway.NewTabDeleteGateway,

			tab_put_usecase.NewTabPutUsecase,
			tab_put_usecase.NewTabPutInteractor,
			tab_put_repository.NewTabPutRepository,
			tab_put_gateway.NewTabPutGateway,

			tab_patch_usecase.NewTabPatchUsecase,
			tab_patch_usecase.NewTabPatchInteractor,
			tab_patch_repository.NewTabPatchRepository,
			tab_patch_gateway.NewTabPatchGateway,

			tab_controller.NewTabController,

			router.NewEchoServer,
			fx.Annotate(
				router.NewRouter,
				fx.ParamTags(`group:"routes"`),
			),
			handler.AsHandler(task_handler.NewTaskHandler),
			handler.AsHandler(board_handler.NewBoardHandler),
			handler.AsHandler(tab_handler.NewTabHandler),
			handler.AsHandler(handler.NewSignUpHandler),
			handler.AsHandler(handler.NewLoginHandler),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
