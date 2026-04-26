package main

import (
	"fmt"
	"os"

	"github.com/vishalyadav0987/task-tracker-cli/interfaces/cli"
	"github.com/vishalyadav0987/task-tracker-cli/internal/infrastructure/persistence/json"
	"github.com/vishalyadav0987/task-tracker-cli/internal/infrastructure/task"
)

func main() {

	// ---------- Setup file path ----------
	filePath := "internal/infrastructure/persistence/json/store.json"

	// ensure file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		file.Close()
	}

	taskRepo := json.NewTaskRepository(filePath)
	taskService := task.NewTaskService(taskRepo)

	handler := cli.NewHandler(taskService)

	handler.Run()

}
