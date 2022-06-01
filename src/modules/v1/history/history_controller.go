package history

import (
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
	result, err := c.repository.FindAll()
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
	}

	helper.ResponseJSON(w, http.StatusOK, result)
}

func (c *controller) GetHistorys(w http.ResponseWriter, r *http.Request) {
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

func (c *controller) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	param, err := strconv.Atoi(params)
	if err != nil {
		fmt.Println("error")
	}

	_, err = c.repository.GetID(param)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Data not found")
		return
	}

	err = c.repository.Delete(param)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Failed delete data history")
		return
	}

	helper.ResponseJSON(w, http.StatusOK, "Success delete history")
}
