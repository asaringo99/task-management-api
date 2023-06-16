package presenter

import (
	usecase "github.com/asaringo99/task_management/internal/application/usecase/board/fetch"
)

type BoardFetchPresenter struct {
	taskFetchUsecaseOutput []usecase.BoardFetchUsecaseOutput
}

type BoardFetchPresenterOutputDto struct {
	Boardid  int    `json:"boardid"`
	Userid   int    `json:"userid"`
	Priority int    `json:"prioriry"`
	Status   string `json:"status"`
}

func (presenter BoardFetchPresenter) Build() []BoardFetchPresenterOutputDto {
	output := presenter.taskFetchUsecaseOutput
	ret := []BoardFetchPresenterOutputDto{}
	for _, res := range output {
		ret = append(ret, convert(res))
	}
	return ret
}

func convert(output usecase.BoardFetchUsecaseOutput) BoardFetchPresenterOutputDto {
	return BoardFetchPresenterOutputDto{
		Boardid:  output.Boardid.ToValue(),
		Userid:   output.Userid.ToValue(),
		Priority: output.Priority.ToValue(),
		Status:   output.Status.ToValue(),
	}
}

func NewBoardAllGetPresenter(output []usecase.BoardFetchUsecaseOutput) BoardFetchPresenter {
	return BoardFetchPresenter{
		taskFetchUsecaseOutput: output,
	}
}
