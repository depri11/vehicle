package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := c.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Send(w)

}

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	res, err := c.service.RegisterUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}

func (c *controller) GetUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	param, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	result, err := c.service.FindByID(param)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Send(w)
}

func (c *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	reqUserId := r.Header.Get("user_id")
	if reqUserId != params {
		http.Error(w, "access danied", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(params)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var inputData UserInput

	json.NewDecoder(r.Body).Decode(&inputData)

	result, err := c.service.UpdateUser(id, &inputData)
	if err != nil {
		http.Error(w, "fail update data", http.StatusBadRequest)
		return
	}

	result.Send(w)
}

func (c *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	// reqUserId := r.Header.Get("user_id")
	id, err := strconv.Atoi(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	res, err := c.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)

}
