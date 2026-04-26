package task

import "context"

type TaskRepository interface {
	Save(ctx context.Context, task *Task) error
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*Task, error)
	GetByStatus(ctx context.Context, status string) ([]*Task, error)
	GetByID(ctx context.Context, id string) (*Task, error)
}

// 🧠 Why context?

// Because:

// DB call slow ho sakta hai
// Timeout cancel karna pad sakta hai
// Production safe code
