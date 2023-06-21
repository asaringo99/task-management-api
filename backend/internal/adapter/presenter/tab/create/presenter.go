package presenter

import (
	usecase "github.com/asaringo99/task_management/internal/application/usecase/tab/create"
)

type TabCreatePresenter struct {
	taskCreateUsecaseOutput usecase.TabCreateUsecaseOutput
}

type TabCreatePresenterOutputDto struct {
	Tabid    int    `json:"tabid"`
	Userid   int    `json:"userid"`
	IsActive bool   `json:"isactive"`
	Title    string `json:"title"`
}

func (presenter TabCreatePresenter) Build() *TabCreatePresenterOutputDto {
	output := convert(presenter.taskCreateUsecaseOutput)
	return &output
}

func convert(output usecase.TabCreateUsecaseOutput) TabCreatePresenterOutputDto {
	return TabCreatePresenterOutputDto{
		Tabid:    output.Tabid.ToValue(),
		Userid:   output.Userid.ToValue(),
		IsActive: output.IsActive,
		Title:    output.Title.ToValue(),
	}
}

func NewTabAllGetPresenter(output usecase.TabCreateUsecaseOutput) TabCreatePresenter {
	return TabCreatePresenter{
		taskCreateUsecaseOutput: output,
	}
}
