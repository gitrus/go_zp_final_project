package main

import (
	"fmt"
	"go_zp_final_project/internal/domain/ports"
	taskservice "go_zp_final_project/internal/domain/services/tasksservice"
	"go_zp_final_project/internal/taskhandlers"
	"go_zp_final_project/pkg/middlewares"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var storage ports.Storage = struct{}{}
	var taskService *taskservice.TaskService = taskservice.NewTaskService(storage)

	var webDir string = getWebDir()
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	// task router
	taskRouter := http.NewServeMux()
	taskRouter.HandleFunc(
		"POST /api/task",
		taskhandlers.CreateTaskHandler(taskService),
	)

	preparedTaskRouter := middlewares.JsonMiddleware(taskRouter)
	http.Handle("/api/task", preparedTaskRouter)

	port := ":7540"
	fmt.Printf("Starting server on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(fmt.Errorf("Failed to start server: %v\n", err))
	}

}

func getWebDir() string {
	thisDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(fmt.Errorf("Cant build path for current dir %e", err))
	}

	var webDirRel string = filepath.Join(thisDir, "..", "..", "web")
	webDir, err := filepath.Abs(webDirRel)
	if err != nil {
		panic(fmt.Errorf("Cant build path for web dir %e", err))
	}
	fmt.Printf("Web dir: %s", thisDir)

	return webDir
}
