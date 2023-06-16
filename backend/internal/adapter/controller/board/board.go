package controller

import (
	usecase_create "github.com/asaringo99/task_management/internal/application/usecase/board/create"
	usecase_delete "github.com/asaringo99/task_management/internal/application/usecase/board/delete"
	usecase_get "github.com/asaringo99/task_management/internal/application/usecase/board/fetch"
	usecase_put "github.com/asaringo99/task_management/internal/application/usecase/board/put"
)

type BoardController struct {
	boardFetchUsecase  usecase_get.BoardFetchUsecaseInputPort
	boardCreateUsecase usecase_create.BoardCreateUsecaseInputPort
	boardDeleteUsecase usecase_delete.BoardDeleteUsecaseInputPort
	boardPutUsecase    usecase_put.BoardPutUsecaseInputPort
}

func NewBoardController(
	boardGetUsecase usecase_get.BoardFetchUsecaseInputPort,
	boardCreateUsecase usecase_create.BoardCreateUsecaseInputPort,
	boardDeleteUsecase usecase_delete.BoardDeleteUsecaseInputPort,
	boardPutUsecase usecase_put.BoardPutUsecaseInputPort,
) *BoardController {
	return &BoardController{
		boardFetchUsecase:  boardGetUsecase,
		boardCreateUsecase: boardCreateUsecase,
		boardDeleteUsecase: boardDeleteUsecase,
		boardPutUsecase:    boardPutUsecase,
	}
}
