package task

import "time"

type Status string

type Task struct {
	id          string
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTask(
	id,
	description string,
	status Status,
) (*Task, error) {
	if description == "" {
		return nil, ErrInvalidDescription
	}
	now := time.Now()

	return &Task{
		id:          id,
		description: description,
		status:      status,
		createdAt:   now,
		updatedAt:   now,
	}, nil
}

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

func (t *Task) ID() string {
	return t.id
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() Status {
	return t.status
}

func (t *Task) MarkDone() {
	t.status = StatusDone
	t.updatedAt = time.Now()
}

func (t *Task) MarkInProgress() {
	t.status = StatusInProgress
	t.updatedAt = time.Now()
}

func (t *Task) UpdateDescription(desc string) error {
	if desc == "" {
		return ErrInvalidDescription
	}
	t.description = desc
	t.updatedAt = time.Now()
	return nil
}
