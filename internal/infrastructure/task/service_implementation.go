package task

import (
	"context"

	"github.com/google/uuid"
	"github.com/vishalyadav0987/task-tracker-cli/internal/application/task/dto"
	domain "github.com/vishalyadav0987/task-tracker-cli/internal/domain/task"
	"github.com/vishalyadav0987/task-tracker-cli/internal/infrastructure/persistence/json"
)

type TaskService struct {
	repo *json.TaskRepository
}

func NewTaskService(repo *json.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

// // ---------------- Add Task ----------------

func (s *TaskService) AddTask(ctx context.Context, input dto.AddTaskInput) error {
	task, err := domain.NewTask(generateID(), input.Description)
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, task)
}

// ---------------- List Tasks ----------------

func (s *TaskService) ListTasks(ctx context.Context) ([]*dto.TaskDTO, error) {
	tasks, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []*dto.TaskDTO

	for _, t := range tasks {
		result = append(result, &dto.TaskDTO{
			ID:          t.ID,
			Description: t.Description,
			Status:      string(t.Status),
		})
	}

	return result, nil
}

// ---------------- Delete Task ----------------

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// ---------------- Update Task ----------------

func (s *TaskService) UpdateTask(ctx context.Context, input dto.UpdateTaskInput) error {
	task, err := s.repo.GetByID(ctx, input.ID)
	if err != nil {
		return err
	}

	if input.Description != "" {
		if err := task.UpdateDescription(input.Description); err != nil {
			return err
		}
	}

	if input.Status != "" {
		if !domain.Status(input.Status).IsValid() {
			return domain.ErrInvalidStatus
		}
		task.UpdateStatus(domain.Status(input.Status))
	}

	return s.repo.Update(ctx, task)
}

// ---------------- Get By Status ----------------

func (s *TaskService) GetTasksByStatus(ctx context.Context, status string) ([]*dto.TaskDTO, error) {
	if !domain.Status(status).IsValid() {
		return nil, domain.ErrInvalidStatus
	}

	tasks, err := s.repo.GetByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	var result []*dto.TaskDTO

	for _, t := range tasks {
		result = append(result, &dto.TaskDTO{
			ID:          t.ID,
			Description: t.Description,
			Status:      string(t.Status),
		})
	}

	return result, nil
}

// ---------------- Get By ID ----------------

func (s *TaskService) GetTasksById(ctx context.Context, id string) (*dto.TaskDTO, error) {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.TaskDTO{
		ID:          task.ID,
		Description: task.Description,
		Status:      string(task.Status),
	}, nil
}

// ---------------- Mark Done ----------------

func (s *TaskService) MarkDone(ctx context.Context, id string) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	task.MarkDone()

	return s.repo.Update(ctx, task)
}

// ---------------- Mark Progess ----------------

func (s *TaskService) MarkProgress(ctx context.Context, id string) error {
	task, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	task.MarkInProgress()

	return s.repo.Update(ctx, task)
}

// ---------------- helper ----------------

func generateID() string {
	return uuid.New().String()
}
