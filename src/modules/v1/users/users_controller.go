package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/helper"
	"github.com/depri11/vehicle/src/interfaces"
	"github.com/gorilla/mux"
)

type controller struct {
	service interfaces.UserService
}

func NewController(service interfaces.UserService) *controller {
	return &controller{service}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Send(w)

}

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	err := helper.ValidationError(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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

	id, err := strconv.Atoi(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var inputData models.User

	json.NewDecoder(r.Body).Decode(&inputData)

	err = helper.ValidationError(inputData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := c.service.UpdateUser(id, &inputData, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Send(w)
}

func (c *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)["id"]

	id, err := strconv.Atoi(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	res, err := c.service.Delete(id, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}

func (c *controller) Query(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sort := r.URL.Query().Get("sort")
	search := r.URL.Query().Get("search")

	lowerSearch := strings.ToLower(search)

	if search != "" {
		res, err := c.service.Search(lowerSearch)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Send(w)
		return
	}

	if sort == "asc" {
		res, err := c.service.Sort(sort)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.Send(w)
		return
	}

	result, err := c.service.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result.Send(w)
}
