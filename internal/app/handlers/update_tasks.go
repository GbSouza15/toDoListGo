package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/GbSouza15/apiToDoGo/internal/app/response"
	"github.com/gorilla/mux"
)

func (h handler) UpdateTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.SendResponse(500, []byte("Error reading the request body"), w)
		return
	}

	var updateTask models.Task

	if err := json.Unmarshal(body, &updateTask); err != nil {
		response.SendResponse(500, []byte("Error decoding JSON"), w)
		return
	}

	_, err = h.DB.Exec("UPDATE tdlist.tasks SET title = $1, description = $2 WHERE id = $3", updateTask.Title, updateTask.Description, taskId)
	if err != nil {
		response.SendResponse(500, []byte("Error updating the task."), w)
		fmt.Println(err)
		return
	}

	response.SendResponse(200, []byte("Task updated successfully"), w)

}
