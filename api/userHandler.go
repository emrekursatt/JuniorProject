package api

import (
	"encoding/json"
	"github.com/emrekursatt/JuniorProject/models"
	"github.com/emrekursatt/JuniorProject/services"
	"net/http"
)

type UserHandler struct {
	Service services.UserService
}

// @Summary Add a new user
// @Description add by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user     body    models.UserEntity     true  "Add User"
// @Success 200 {object}  models.UserEntity
// @Router /user/create [post]
func (h UserHandler) Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.UserEntity

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
	}

	result, err := h.Service.Insert(user)
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

// @Summary Add a new user
// @Description add by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user     body    models.UserEntity     true  "Delete User"
// @Success 200 {object}  models.UserEntity
// @Router /user/delete [DELETE]
func (h UserHandler) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.UserEntity

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
	}

	result, err := h.Service.Delete(user)
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

// @Summary Add a new user
// @Description add by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user     body    models.UserEntity     true  "Update User"
// @Success 200 {object} models.UserEntity
// @Router /user/update [PUT]
func (h UserHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user models.UserEntity

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return

		}
	}

	result, err := h.Service.UpdatePassword(user)
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

// @Summary Add a new user
// @Description add by json user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object}  models.UserEntity
// @Router /user/getAll [GET]
func (h UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var users []models.UserEntity

	users, err := h.Service.GetAllUsers()
	if err != nil {
		errMsg := "Internal Server Error"
		if err != nil {
			errMsg = err.Error()
		}
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
	w.WriteHeader(http.StatusOK)
}
