package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/vishalyadav0987/task-tracker-cli/internal/application/task/dto"
	app "github.com/vishalyadav0987/task-tracker-cli/internal/infrastructure/task"
)

type Handler struct {
	service *app.TaskService
}

func NewHandler(service *app.TaskService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Run() {
	cmd, err := Parse(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ctx := context.Background()

	switch cmd.Name {

	case "add":
		if len(cmd.Args) < 1 {
			fmt.Println("Usage: task add \"description\"")
			return
		}

		input := dto.AddTaskInput{
			Description: cmd.Args[0],
		}

		err := h.service.AddTask(ctx, input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task added successfully")

	case "list":
		tasks, err := h.service.ListTasks(ctx)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return
		}

		for _, t := range tasks {
			fmt.Printf("ID: %s | %s | %s\n", t.ID, t.Description, t.Status)
		}

	case "update":
		if len(cmd.Args) < 2 {
			fmt.Println("Usage: task update <id> \"new description\"")
			return
		}

		desc := cmd.Args[1]

		input := dto.UpdateTaskInput{
			ID:          cmd.Args[0],
			Description: desc,
		}

		err := h.service.UpdateTask(ctx, input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task updated successfully")

	case "delete":
		if len(cmd.Args) < 1 {
			fmt.Println("Usage: task delete <id>")
			return
		}

		err := h.service.DeleteTask(ctx, cmd.Args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task deleted")

	case "mark-in-progress":
		if len(cmd.Args) < 1 {
			fmt.Println("Usage: task in-progress <id>")
			return
		}

		err := h.service.MarkProgress(ctx, cmd.Args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task marked as in-progress")

	case "mark-done":
		if len(cmd.Args) < 1 {
			fmt.Println("Usage: task done <id>")
			return
		}

		err := h.service.MarkDone(ctx, cmd.Args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Task marked as done")

	case "status":
		if len(cmd.Args) < 1 {
			fmt.Println("Usage: task status <todo|in-progress|done>")
			return
		}

		status := cmd.Args[0]
		switch status {
		case "done":
			status = "mark-done"
		case "in-progess":
			status = "mark-in-progress"
		case "todo":
			status = "mark-todo"
		}

		tasks, err := h.service.GetTasksByStatus(ctx, status)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, t := range tasks {
			fmt.Printf("ID: %s | %s | %s\n", t.ID, t.Description, t.Status)
		}
	case "task":
		if len(cmd.Args) < 1 {
			fmt.Println("Usage: task get by Id <id>")
			return
		}

		task, err := h.service.GetTasksById(ctx, cmd.Args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("ID: %s | %s | %s\n", task.ID, task.Description, task.Status)

	default:
		fmt.Println("Unknown command")
	}
}
