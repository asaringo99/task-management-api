package presenter

import (
	usecase "github.com/asaringo99/task_management/internal/application/usecase/task/fetch"
)

type TaskFetchPresenter struct {
	taskFetchUsecaseOutput []usecase.TaskFetchUsecaseOutput
}

type TaskFetchPresenterOutputDto struct {
	Taskid   int    `json:"taskid"`
	Userid   int    `json:"userid"`
	Boardid  int    `json:"boardid"`
	Priority int    `json:"priority"`
	Contents string `json:"contents"`
}

func (presenter TaskFetchPresenter) Build() []TaskFetchPresenterOutputDto {
	output := presenter.taskFetchUsecaseOutput
	ret := []TaskFetchPresenterOutputDto{}
	for _, res := range output {
		ret = append(ret, convert(res))
	}
	return ret
}

func convert(output usecase.TaskFetchUsecaseOutput) TaskFetchPresenterOutputDto {
	return TaskFetchPresenterOutputDto{
		Taskid:   output.Taskid.ToValue(),
		Userid:   output.Userid.ToValue(),
		Boardid:  output.Boardid.ToValue(),
		Priority: output.Priority.ToValue(),
		Contents: output.Contents.ToValue(),
	}
}

func NewTaskAllGetPresenter(output []usecase.TaskFetchUsecaseOutput) TaskFetchPresenter {
	return TaskFetchPresenter{
		taskFetchUsecaseOutput: output,
	}
}
