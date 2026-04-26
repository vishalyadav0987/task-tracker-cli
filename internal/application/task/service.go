package taskService

import (
	"context"

	"github.com/vishalyadav0987/task-tracker-cli/internal/application/task/dto"
)

type TaskService interface {
	AddTask(ctx context.Context, input dto.AddTaskInput) error
	ListTasks(ctx context.Context) ([]*dto.TaskDTO, error)
	DeleteTask(ctx context.Context, id string) error
	MarkDone(ctx context.Context, id string) error
	UpdateTask(ctx context.Context, input dto.UpdateTaskInput) error
	GetTasksByStatus(ctx context.Context, status string) ([]*dto.TaskDTO, error)
	GetTasksById(ctx context.Context, id string) ([]*dto.TaskDTO, error)
}
