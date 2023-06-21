package usecase

import (
	"reflect"
	"testing"

	repository "github.com/asaringo99/task_management/internal/application/repository/tab/fetch"
	domain "github.com/asaringo99/task_management/internal/domain/valueobject"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name   string
		output repository.TabFetchRepositoryOutput
		want   TabFetchUsecaseOutput
	}{
		{
			name: "no diff1",
			output: repository.TabFetchRepositoryOutput{
				Tabid:    domain.NewId(1),
				Userid:   domain.NewUserid(1),
				IsActive: true,
				Title:    domain.NewTitle("test"),
			},
			want: TabFetchUsecaseOutput{
				Tabid:    domain.NewId(1),
				Userid:   domain.NewUserid(1),
				IsActive: true,
				Title:    domain.NewTitle("test"),
			},
		},
		// TODO: Add Tests
	}

	for _, tc := range tests {
		got := convert(tc.output)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("case: %s, expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
