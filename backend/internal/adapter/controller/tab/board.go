package controller

import (
	usecase_create "github.com/asaringo99/task_management/internal/application/usecase/tab/create"
	usecase_delete "github.com/asaringo99/task_management/internal/application/usecase/tab/delete"
	usecase_get "github.com/asaringo99/task_management/internal/application/usecase/tab/fetch"
	usecase_patch "github.com/asaringo99/task_management/internal/application/usecase/tab/patch"
	usecase_put "github.com/asaringo99/task_management/internal/application/usecase/tab/put"
)

type TabController struct {
	tabFetchUsecase  usecase_get.TabFetchUsecaseInputPort
	tabCreateUsecase usecase_create.TabCreateUsecaseInputPort
	tabDeleteUsecase usecase_delete.TabDeleteUsecaseInputPort
	tabPutUsecase    usecase_put.TabPutUsecaseInputPort
	tabPatchUsecase  usecase_patch.TabPatchUsecaseInputPort
}

func NewTabController(
	tabGetUsecase usecase_get.TabFetchUsecaseInputPort,
	tabCreateUsecase usecase_create.TabCreateUsecaseInputPort,
	tabDeleteUsecase usecase_delete.TabDeleteUsecaseInputPort,
	tabPutUsecase usecase_put.TabPutUsecaseInputPort,
	tabPatchUsecase usecase_patch.TabPatchUsecaseInputPort,
) *TabController {
	return &TabController{
		tabFetchUsecase:  tabGetUsecase,
		tabCreateUsecase: tabCreateUsecase,
		tabDeleteUsecase: tabDeleteUsecase,
		tabPutUsecase:    tabPutUsecase,
		tabPatchUsecase:  tabPatchUsecase,
	}
}
