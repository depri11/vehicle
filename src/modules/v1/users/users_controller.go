package users

import (
	"encoding/json"
	"net/http"

	"github.com/depri11/vehicle/src/helper"
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
		res := helper.ResponseJSON("Failed to get all data users", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	res := helper.ResponseJSON("List of Users", 200, "success", result)
	json.NewEncoder(w).Encode(res)

}

func (c *controller) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

	json.NewDecoder(r.Body).Decode(&user)
	res, err := c.service.RegisterUser(&user)
	if err != nil {
		res := helper.ResponseJSON("Failed to register user", http.StatusNotFound, "error", err.Error())
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(res)
}

// func (c *controller) Create(w http.ResponseWriter, r *http.Request) {
// 	var data User
// 	json.NewDecoder(r.Body).Decode(&data)

// 	result, err := c.service.RegisterUser(&data)
// 	if err != nil {
// 		res := helper.ResponseJSON("Failed create data user", http.StatusBadRequest, "error", err.Error())
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	res := helper.ResponseJSON("List of Users", 200, "success", result)
// 	json.NewEncoder(w).Encode(res)
// }

// func (c *controller) GetUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)["id"]
// 	param, err := strconv.Atoi(params)
// 	if err != nil {
// 		fmt.Println("error")
// 	}

// 	result, err := c.repository.GetUserID(param)
// 	if err != nil {
// 		res := helper.ResponseJSON("Failed get user", http.StatusNotFound, "error", err.Error())
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	res := helper.ResponseJSON("List data User", 200, "success", result)
// 	json.NewEncoder(w).Encode(res)
// }

// func (c *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)["id"]
// 	id, err := strconv.Atoi(params)
// 	if err != nil {
// 		res := helper.ResponseJSON("Internal Server Error", http.StatusInternalServerError, "error", err.Error())
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	user, err := c.repository.GetUserID(id)
// 	if err != nil {
// 		res := helper.ResponseJSON("Failed get User", http.StatusNotFound, "error", err.Error())
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	json.NewDecoder(r.Body).Decode(&user)

// 	result, err := c.repository.Update(user)
// 	if err != nil {
// 		res := helper.ResponseJSON("Failed update data user", http.StatusBadRequest, "error", err.Error())
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	res := helper.ResponseJSON("Success update data User", 200, "success", result)
// 	json.NewEncoder(w).Encode(res)
// }

// func (c *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)["id"]
// 	id, err := strconv.Atoi(params)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	_, err = c.repository.GetUserID(id)
// 	if err != nil {
// 		res := helper.ResponseJSON("Failed get user", http.StatusNotFound, "error", err.Error())
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	err = c.repository.Delete(id)
// 	if err != nil {
// 		res := helper.ResponseJSON("Failed delete user", http.StatusNotFound, "error", err.Error())
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	res := helper.ResponseJSON("Successfully delete user", 200, "success", nil)
// 	json.NewEncoder(w).Encode(res)

// }
