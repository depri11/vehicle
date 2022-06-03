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
	repository Repository
}

func NewController(repository Repository) *controller {
	return &controller{repository}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := c.repository.FindAll()
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get Vehicle", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}
	res := helper.ResponseJSON(w, "List of Vehicle", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)

	result, err := c.repository.Save(&vehicle)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed create Vehicle", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "Successfully created Vehicle", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)

}

func (c *controller) GetVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	param, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	result, err := c.repository.GetID(param)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get Vehicle", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "List of Vehicle", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)

}

func (c *controller) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	vehicle, err := c.repository.GetID(id)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get Vehicle", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewDecoder(r.Body).Decode(&vehicle)

	result, err := c.repository.Update(vehicle)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed update Vehicle", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "Successfully updated Vhicle", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)

}

func (c *controller) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, err := strconv.Atoi(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = c.repository.GetID(id)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get Vehicle", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	err = c.repository.Delete(id)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed deleted Vehicle", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "Successfully deleted Vehicle", http.StatusOK, "success", nil)
	json.NewEncoder(w).Encode(res)
}

func (c *controller) PopularVehicle(w http.ResponseWriter, r *http.Request) {
	result, err := c.repository.Popular()
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get popular Vehicle", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "List popular Vehicle", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)
}

func (c *controller) QuerySort(w http.ResponseWriter, r *http.Request) {
	sort := r.URL.Query().Get("sort")
	search := r.URL.Query().Get("search")

	string := strings.ToLower(search)

	if search != "" {
		result, err := c.repository.Search(string)
		if err != nil {
			res := helper.ResponseJSON(w, "Internal Server Error", http.StatusInternalServerError, "error", err.Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res := helper.ResponseJSON(w, "List data Search", http.StatusOK, "success", result)
		json.NewEncoder(w).Encode(res)
		return
	}

	if sort == "asc" {
		result, err := c.repository.Query(sort)
		if err != nil {
			res := helper.ResponseJSON(w, "Internal Server Error", http.StatusInternalServerError, "error", err.Error())
			json.NewEncoder(w).Encode(res)
			return
		}

		res := helper.ResponseJSON(w, "List data Search", http.StatusOK, "success", result)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := c.repository.FindAll()
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get Vehicles", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "List of Vehicle", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)
}
