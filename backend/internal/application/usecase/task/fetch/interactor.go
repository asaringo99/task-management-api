package usecase

import (
	repository "github.com/asaringo99/task_management/internal/application/repository/task/fetch"
	"github.com/asaringo99/task_management/internal/application/usecase/task/fetch/condition"
)

type TaskFetchInteractor struct {
	repository repository.TaskFetchRepositoryInterface
}

func NewTaskFetchInteractor(repository repository.TaskFetchRepositoryInterface) *TaskFetchInteractor {
	return &TaskFetchInteractor{
		repository: repository,
	}
}

func (interactor *TaskFetchInteractor) find(input condition.TaskFetchCondition) ([]TaskFetchUsecaseOutput, error) {
	response, err := interactor.repository.Find(input)
	if err != nil {
		return nil, err
	}
	ret := []TaskFetchUsecaseOutput{}
	for _, res := range response {
		ret = append(ret, convert(res))
	}
	return ret, nil
}

func convert(value repository.TaskFetchRepositoryOutput) TaskFetchUsecaseOutput {
	return TaskFetchUsecaseOutput{
		Taskid:   value.Taskid,
		Userid:   value.Userid,
		Status:   value.Status,
		Contents: value.Contents,
		Priority: value.Priority,
	}
}
