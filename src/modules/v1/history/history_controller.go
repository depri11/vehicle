package history

import (
	"encoding/json"
	"fmt"
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
		res := helper.ResponseJSON(w, "Failed to get Historys", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "List of Historys", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)
}

func (c *controller) GetHistorys(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	param, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	result, err := c.repository.GetID(param)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get Historys", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "List user data", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)
}

func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
	var data Historys
	json.NewDecoder(r.Body).Decode(&data)

	result, err := c.repository.Save(&data)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed create Historys", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "Successfully create History", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)
}

func (c *controller) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	param, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	_, err = c.repository.GetID(param)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get History", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	err = c.repository.Delete(param)
	if err != nil {
		res := helper.ResponseJSON(w, "Failed delete History", http.StatusBadRequest, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "Successfully deleted History", http.StatusOK, "success", nil)
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

		res := helper.ResponseJSON(w, "List search data", http.StatusOK, "success", result)
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

		res := helper.ResponseJSON(w, "List sort data", http.StatusOK, "success", result)
		json.NewEncoder(w).Encode(res)
		return
	}

	result, err := c.repository.FindAll()
	if err != nil {
		res := helper.ResponseJSON(w, "Failed get Historys", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON(w, "List of Historys", http.StatusOK, "success", result)
	json.NewEncoder(w).Encode(res)
}
