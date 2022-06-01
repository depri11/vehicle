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
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(data)
	helper.ResponseJSON(w, http.StatusOK, result)

}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	var data User
	json.NewDecoder(r.Body).Decode(&data)

	result, err := c.repository.Save(&data)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// json.NewEncoder(w).Encode(&result)
	helper.ResponseJSON(w, http.StatusOK, result)
}

func (c *controller) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	param, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	result, err := c.repository.GetUserID(param)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "User not found")
		return
	}

	// json.NewEncoder(w).Encode(&result)
	helper.ResponseJSON(w, http.StatusOK, result)
}

func (c *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	user, err := c.repository.GetUserID(id)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "User not found")
		return
	}

	json.NewDecoder(r.Body).Decode(&user)

	result, err := c.repository.Update(user)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Failed update data user")
		return
	}

	// json.NewEncoder(w).Encode(result)
	helper.ResponseJSON(w, http.StatusOK, result)
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
		helper.ResponseError(w, http.StatusBadRequest, "User not found")
		return
	}

	err = c.repository.Delete(id)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Failed delete data user")
		return
	}

	helper.ResponseJSON(w, http.StatusOK, "Success delete user")

}
