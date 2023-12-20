package api

import (
	"encoding/json"
	"github.com/emrekursatt/JuniorProject/models"
	"github.com/emrekursatt/JuniorProject/services"
	"net/http"
)

type TaskHandler struct {
	Service services.TaskService
}

// @Summary Add a new task
// @Description add by json task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   task     body    models.TaskEntity     true  "Add Task"
// @Success 200 {object}  models.TaskEntity
// @Router /task/create [post]
func (h TaskHandler) Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var task models.TaskEntity

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
	}

	result, err := h.Service.Insert(task)
	if err != nil || result.Status == false {
		errMsg := "Internal Server Error"
		if err != nil {
			errMsg = err.Error()
		}
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Add a new task
// @Description add by json task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   task     body    models.TaskEntity     true  "Delete Task"
// @Success 200 {object}  models.TaskEntity
// @Router /task/delete [DELETE]
func (h TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var task models.TaskEntity

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
	}

	result, err := h.Service.Delete(task)
	if err != nil || result.Status == false {
		errMsg := "Internal Server Error"
		if err != nil {
			errMsg = err.Error()
		}
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Add a new task
// @Description add by json task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param   task     body    models.TaskEntity     true  "Update Task"
// @Success 200 {object}  models.TaskEntity
// @Router /task/update [PUT]
func (h TaskHandler) Update(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var task models.TaskEntity

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
	}

	result, err := h.Service.Update(task)
	if err != nil || result.Status == false {
		errMsg := "Internal Server Error"
		if err != nil {
			errMsg = err.Error()
		}
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @Summary Add a new task
// @Description add by json task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {object}  []models.TaskEntity
// @Router /task/getAll [GET]
func (h TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	tasks, err := h.Service.GetAllTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}
