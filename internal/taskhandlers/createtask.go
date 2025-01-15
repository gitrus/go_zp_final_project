package taskhandlers

import (
	"encoding/json"
	"fmt"
	"go_zp_final_project/internal/domain/entities"
	"go_zp_final_project/internal/domain/services/repeatrules"
	taskservice "go_zp_final_project/internal/domain/services/tasksservice"
	"go_zp_final_project/pkg/tasksschemas"
	"net/http"
	"time"
)

func CreateTaskHandler(taskService *taskservice.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var task tasksschemas.TaskSchema
		decoder := json.NewDecoder(r.Body)
		if err = decoder.Decode(&task); err != nil {
			http.Error(w, "Bad request: invalid JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received Task: %+v\n", task)

		var rule entities.RepeatRule
		rule, err = repeatrules.ParseRule(task.Repeat)

		var netExec time.Time
		netExec, err = rule.NextExecution(time.Now())
		fmt.Print("Prased repeat rule: ", rule, netExec)

		var domainTask entities.Task
		if err := taskService.CreateTask(domainTask); err != nil {
			http.Error(w, fmt.Sprintf("Failed to create task: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\":\"Task received successfully\"}"))
	}
}
