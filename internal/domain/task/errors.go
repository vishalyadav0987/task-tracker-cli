package task

import "errors"

var (
	ErrTaskNotFound       = errors.New("task not found")
	ErrInvalidDescription = errors.New("invalid task description")
	ErrInvalidStatus      = errors.New("invalid task status")
)
