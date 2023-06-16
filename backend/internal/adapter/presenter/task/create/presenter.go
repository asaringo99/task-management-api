package presenter

import (
	usecase "github.com/asaringo99/task_management/internal/application/usecase/task/create"
)

type TaskCreatePresenter struct {
	taskCreateUsecaseOutput usecase.TaskCreateUsecaseOutput
}

type TaskCreatePresenterOutputDto struct {
	Taskid   int    `json:"taskid"`
	Userid   int    `json:"userid"`
	Boardid  int    `json:"boardid"`
	Priority int    `json:"priority"`
	Contents string `json:"contents"`
}

func (presenter TaskCreatePresenter) Build() *TaskCreatePresenterOutputDto {
	output := presenter.taskCreateUsecaseOutput
	ret := convert(output)
	return &ret
}

func convert(output usecase.TaskCreateUsecaseOutput) TaskCreatePresenterOutputDto {
	return TaskCreatePresenterOutputDto{
		Taskid:   output.Taskid.ToValue(),
		Userid:   output.Userid.ToValue(),
		Boardid:  output.Boardid.ToValue(),
		Priority: output.Priority.ToValue(),
		Contents: output.Contents.ToValue(),
	}
}

func NewTaskCreatePresenter(output usecase.TaskCreateUsecaseOutput) *TaskCreatePresenter {
	return &TaskCreatePresenter{
		taskCreateUsecaseOutput: output,
	}
}
