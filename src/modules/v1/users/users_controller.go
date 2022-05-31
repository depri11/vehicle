package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/depri11/vehicle/src/helper"
	"github.com/gorilla/mux"
)

type controller struct {
	repository *repository
}

func NewController(repository *repository) *controller {
	return &controller{repository}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := c.repository.FindAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data User
	json.NewDecoder(r.Body).Decode(&data)

	result, err := c.repository.Save(&data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (c *controller) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)["id"]
	param, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	result, err := c.repository.GetUserID(param)
	if err != nil {
		fmt.Fprint(w, errors.New("User not found"))
	}

	json.NewEncoder(w).Encode(&result)

}

func (c *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	user, err := c.repository.GetUserID(id)
	if err != nil {
		fmt.Fprint(w, errors.New("User not found"))
	}

	json.NewDecoder(r.Body).Decode(&user)

	result, err := c.repository.Update(user)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(result)

}

func (c *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	err = c.repository.Delete(id)
	if err != nil {
		errors.New("failed delete user")
	}

	helper.RespondJSON(w, http.StatusOK, "Delete user successfully")
}
