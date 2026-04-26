package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/vishalyadav0987/task-tracker-cli/internal/application/task/dto"
	app "github.com/vishalyadav0987/task-tracker-cli/internal/infrastructure/task"
)

type Handler struct {
	service *app.TaskService
}

func NewHandler(service *app.TaskService) *Handler {
	return &Handler{service: service}
}

func RenderTasks(tasks []*dto.TaskDTO) {
	table := tablewriter.NewWriter(os.Stdout)

	table.Header([]string{"ID", "Description", "Status"})

	for _, t := range tasks {

		var status string

		switch t.Status {
		case "done":
			status = color.New(color.FgGreen).Sprint("✔ Done")
		case "in-progress":
			status = color.New(color.FgYellow).Sprint("⏳ In Progress")
		case "todo":
			status = color.New(color.FgBlue).Sprint("📌 Todo")
		default:
			status = t.Status
		}

		table.Append([]string{
			t.ID,
			t.Description,
			status,
		})
	}

	table.Render()
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
			PrintInfo("Usage: task add \"description\"")
			return
		}

		input := dto.AddTaskInput{
			Description: cmd.Args[0],
		}

		err := h.service.AddTask(ctx, input)
		if err != nil {
			PrintError(err)
			return
		}

		PrintSuccess("Task added successfully")

	case "list":
		tasks, err := h.service.ListTasks(ctx)
		if err != nil {
			PrintError(err)
			return
		}

		if len(tasks) == 0 {
			PrintWarning("No tasks found")
			return
		}

		// for _, t := range tasks {
		// 	fmt.Printf("ID: %s | %s | %s\n", t.ID, t.Description, t.Status)
		// }

		RenderTasks(tasks)

	case "update":
		if len(cmd.Args) < 2 {
			PrintInfo("Usage: task update <id> \"new description\"")
			return
		}

		desc := cmd.Args[1]

		input := dto.UpdateTaskInput{
			ID:          cmd.Args[0],
			Description: desc,
		}

		err := h.service.UpdateTask(ctx, input)
		if err != nil {
			PrintError(err)
			return
		}

		PrintSuccess("Task updated successfully")

	case "delete":
		if len(cmd.Args) < 1 {
			PrintInfo("Usage: task delete <id>")
			return
		}

		err := h.service.DeleteTask(ctx, cmd.Args[0])
		if err != nil {
			PrintError(err)
			return
		}

		PrintSuccess("Task deleted")

	case "mark-in-progress":
		if len(cmd.Args) < 1 {
			PrintInfo("Usage: task in-progress <id>")
			return
		}

		err := h.service.MarkProgress(ctx, cmd.Args[0])
		if err != nil {
			PrintError(err)
			return
		}

		PrintSuccess("Task marked as in-progress")

	case "mark-done":
		if len(cmd.Args) < 1 {
			PrintInfo("Usage: task done <id>")
			return
		}

		err := h.service.MarkDone(ctx, cmd.Args[0])
		if err != nil {
			PrintError(err)
			return
		}

		PrintSuccess("Task marked as done")

	case "status":
		if len(cmd.Args) < 1 {
			PrintInfo("Usage: task status <todo|in-progress|done>")
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
			PrintError(err)
			return
		}

		// for _, t := range tasks {
		// 	fmt.Printf("ID: %s | %s | %s\n", t.ID, t.Description, t.Status)
		// }
		RenderTasks(tasks)

	case "task":
		if len(cmd.Args) < 1 {
			PrintInfo("Usage: task get by Id <id>")
			return
		}

		task, err := h.service.GetTasksById(ctx, cmd.Args[0])
		if err != nil {
			PrintError(err)
			return
		}

		// fmt.Printf("ID: %s | %s | %s\n", task.ID, task.Description, task.Status)
		RenderTasks([]*dto.TaskDTO{
			task,
		})

	default:
		PrintWarning("Unknown command")
	}
}
