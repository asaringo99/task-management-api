package presenter

import (
	usecase "github.com/asaringo99/task_management/internal/application/usecase/board/create"
)

type BoardCreatePresenter struct {
	taskCreateUsecaseOutput usecase.BoardCreateUsecaseOutput
}

type BoardCreatePresenterOutputDto struct {
	Boardid  int    `json:"boardid"`
	Userid   int    `json:"userid"`
	Tabid    int    `json:"tabid"`
	Priority int    `json:"priority"`
	Status   string `json:"status"`
}

func (presenter BoardCreatePresenter) Build() *BoardCreatePresenterOutputDto {
	output := convert(presenter.taskCreateUsecaseOutput)
	return &output
}

func convert(output usecase.BoardCreateUsecaseOutput) BoardCreatePresenterOutputDto {
	return BoardCreatePresenterOutputDto{
		Boardid:  output.Boardid.ToValue(),
		Userid:   output.Userid.ToValue(),
		Tabid:    output.Tabid.ToValue(),
		Priority: output.Priority.ToValue(),
		Status:   output.Status.ToValue(),
	}
}

func NewBoardAllGetPresenter(output usecase.BoardCreateUsecaseOutput) BoardCreatePresenter {
	return BoardCreatePresenter{
		taskCreateUsecaseOutput: output,
	}
}
