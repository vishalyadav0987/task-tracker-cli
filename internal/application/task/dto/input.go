package dto

type AddTaskInput struct {
	Description string
}

type UpdateTaskInput struct {
	ID          string
	Description string
	Status      string
}
