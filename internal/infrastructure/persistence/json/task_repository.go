package json

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/vishalyadav0987/task-tracker-cli/internal/domain/task"
)

type TaskRepository struct {
	filePath string
	mu       sync.Mutex
}

func NewTaskRepository(filePath string) *TaskRepository {
	return &TaskRepository{
		filePath: filePath,
	}
}

// ---------- helper functions ----------

func (r *TaskRepository) load() ([]*task.Task, error) {
	file, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []*task.Task{}, nil
		}
		return nil, err
	}

	var tasks []*task.Task
	if len(file) == 0 {
		return []*task.Task{}, nil
	}

	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) saveAll(tasks []*task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0644)
}

// ---------- interface implementation ----------

func (r *TaskRepository) Save(ctx context.Context, t *task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.load()
	if err != nil {
		return err
	}

	tasks = append(tasks, t)

	return r.saveAll(tasks)
}

func (r *TaskRepository) Update(ctx context.Context, updated *task.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.load()
	if err != nil {
		return err
	}

	found := false

	for i, t := range tasks {
		if t.ID == updated.ID {
			tasks[i] = updated
			found = true
			break
		}
	}

	if !found {
		return task.ErrTaskNotFound
	}

	return r.saveAll(tasks)
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.load()
	if err != nil {
		return err
	}

	newTasks := make([]*task.Task, 0)

	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		return task.ErrTaskNotFound
	}

	return r.saveAll(newTasks)
}

func (r *TaskRepository) GetAll(ctx context.Context) ([]*task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.load()
}

func (r *TaskRepository) GetByStatus(ctx context.Context, status string) ([]*task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.load()
	if err != nil {
		return nil, err
	}

	var result []*task.Task

	for _, t := range tasks {
		if string(t.Status) == status {
			result = append(result, t)
		}
	}

	return result, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id string) (*task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.load()
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}

	return nil, task.ErrTaskNotFound
}
