package users

import (
	"encoding/json"
	"fmt"
	"log"
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
	result, err := c.repository.FindAll()
	if err != nil {
		res := helper.ResponseJSON(w, "Failed to get all data users", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "List of Users", 200, "success", result)
	json.NewEncoder(w).Encode(res)

}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	var data User
	json.NewDecoder(r.Body).Decode(&data)

	result, err := c.repository.Save(&data)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed create data user", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "List of Users", 200, "success", result)
	json.NewEncoder(w).Encode(res)
}

func (c *controller) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	param, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	result, err := c.repository.GetUserID(param)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get user", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "List data User", 200, "success", result)
	json.NewEncoder(w).Encode(res)
}

func (c *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		res := helper.ResponseJSON(w, "Internal Server Error", http.StatusInternalServerError, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	user, err := c.repository.GetUserID(id)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get User", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)

	result, err := c.repository.Update(user)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed update data user", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "Success update data User", 200, "success", result)
	json.NewEncoder(w).Encode(res)
}

func (c *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = c.repository.GetUserID(id)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get user", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	err = c.repository.Delete(id)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed delete user", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "Successfully delete user", 200, "success", nil)
	json.NewEncoder(w).Encode(res)

}
