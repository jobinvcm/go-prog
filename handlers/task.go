package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RequestData struct {
	TaskName        string `json:"name"`
	TaskDescription string `json:"description"`
}

type Task struct {
	TaskName        string
	TaskDescription string
}

var tasks []Task

func TaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "unalbe to read request Body", http.StatusBadRequest)
			return
		}

		var data RequestData

		err = json.Unmarshal(body, &data)

		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		tasks = append(tasks, Task{TaskName: data.TaskName, TaskDescription: data.TaskDescription})

		fmt.Fprintf(w, "currently we have: %d tasks", len(tasks))
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}

}
