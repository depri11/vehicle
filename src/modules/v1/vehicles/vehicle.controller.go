package vehicle

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/depri11/vehicle/src/helper"
	"github.com/gorilla/mux"
)

type controller struct {
	service Service
}

func NewController(service Service) *controller {
	return &controller{service}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.FindAll()
	if err != nil {
		res := helper.ResponseJSON("Failed get Vehicle", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (c *controller) GetVehicle(w http.ResponseWriter, r *http.Request) {
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

	json.NewEncoder(w).Encode(result)

}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)

	res, err := c.service.Create(&vehicle)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)

}

func (c *controller) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var inputData VehicleInput

	json.NewDecoder(r.Body).Decode(&inputData)

	result, err := c.service.Update(id, &inputData)
	if err != nil {
		http.Error(w, "fail update data", http.StatusBadRequest)
		return
	}

	result.Send(w)
}

func (c *controller) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]

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

func (c *controller) PopularVehicle(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.Popular()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res.Send(w)
}

func (c *controller) Query(w http.ResponseWriter, r *http.Request) {
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
