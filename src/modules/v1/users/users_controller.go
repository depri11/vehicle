package users

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type controller struct {
	repository *repository
}

func NewController(repository *repository) *controller {
	return &controller{repository}
}

func (c *controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := c.repository.FindAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(users)
}
