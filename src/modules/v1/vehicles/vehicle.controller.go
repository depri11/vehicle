package vehicle

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
	}

	helper.ResponseJSON(w, http.StatusOK, result)
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)

	result, err := c.repository.Save(&vehicle)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	// json.NewEncoder(w).Encode(&result)
	helper.ResponseJSON(w, http.StatusOK, result)

}

func (c *controller) GetVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	param, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	result, err := c.repository.GetID(param)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Data not found")
		return
	}

	// json.NewEncoder(w).Encode(&result)
	helper.ResponseJSON(w, http.StatusOK, result)
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
		helper.ResponseError(w, http.StatusBadRequest, "Vehicle not found")
		return
	}

	json.NewDecoder(r.Body).Decode(&vehicle)

	result, err := c.repository.Update(vehicle)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Failed update data vehicle")
		return
	}

	// json.NewEncoder(w).Encode(result)
	helper.ResponseJSON(w, http.StatusOK, result)

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
		helper.ResponseError(w, http.StatusBadRequest, "Vehicle not found")
		return
	}

	err = c.repository.Delete(id)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Failed delete data vehicle")
		return
	}

	helper.ResponseJSON(w, http.StatusOK, "Success delete data")

}

func (c *controller) PopularVehicle(w http.ResponseWriter, r *http.Request) {
	result, err := c.repository.Popular()
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
	}

	helper.ResponseJSON(w, http.StatusOK, result)
}
