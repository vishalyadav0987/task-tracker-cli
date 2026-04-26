package task

import (
	"fmt"
	"time"
)

type Status string

const (
	StatusTodo       Status = "mark-todo"
	StatusInProgress Status = "mark-in-progress"
	StatusDone       Status = "mark-done"
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(id, description string) (*Task, error) {
	if description == "" {
		return nil, ErrInvalidDescription
	}

	now := time.Now()
	fmt.Println("Generated ID:", id)

	return &Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (t *Task) MarkDone() {
	t.Status = StatusDone
	t.UpdatedAt = time.Now()
}

func (t *Task) MarkProgess() {
	t.Status = StatusInProgress
	t.UpdatedAt = time.Now()
}

func (t *Task) MarkInProgress() {
	t.Status = StatusInProgress
	t.UpdatedAt = time.Now()
}

func (t *Task) UpdateDescription(desc string) error {
	if desc == "" {
		return ErrInvalidDescription
	}
	t.Description = desc
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) UpdateStatus(status Status) error {
	if !status.IsValid() {
		return ErrInvalidStatus
	}
	t.Status = status
	t.UpdatedAt = time.Now()
	return nil
}

func (s Status) IsValid() bool {
	switch s {
	case StatusTodo, StatusInProgress, StatusDone:
		return true
	default:
		return false
	}
}
