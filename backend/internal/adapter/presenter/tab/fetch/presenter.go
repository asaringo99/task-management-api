package presenter

import (
	usecase "github.com/asaringo99/task_management/internal/application/usecase/tab/fetch"
)

type TabFetchPresenter struct {
	taskFetchUsecaseOutput []usecase.TabFetchUsecaseOutput
}

type TabFetchPresenterOutputDto struct {
	Tabid    int    `json:"tabid"`
	Userid   int    `json:"userid"`
	IsActive bool   `json:"isactive"`
	Title    string `json:"title"`
}

func (presenter TabFetchPresenter) Build() []TabFetchPresenterOutputDto {
	output := presenter.taskFetchUsecaseOutput
	ret := []TabFetchPresenterOutputDto{}
	for _, res := range output {
		ret = append(ret, convert(res))
	}
	return ret
}

func convert(output usecase.TabFetchUsecaseOutput) TabFetchPresenterOutputDto {
	return TabFetchPresenterOutputDto{
		Tabid:    output.Tabid.ToValue(),
		Userid:   output.Userid.ToValue(),
		IsActive: output.IsActive,
		Title:    output.Title.ToValue(),
	}
}

func NewTabAllGetPresenter(output []usecase.TabFetchUsecaseOutput) TabFetchPresenter {
	return TabFetchPresenter{
		taskFetchUsecaseOutput: output,
	}
}
