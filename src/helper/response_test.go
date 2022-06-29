package helper

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseJSON(t *testing.T) {
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	user := User{ID: 1, Name: "John"}
	data, err := json.Marshal(user)
	assert.Nil(t, err)

	response := ResponseJSON("success", 200, "success", string(data))
	assert.Equal(t, "success", response.Meta.Message)
	assert.Equal(t, 200, response.Meta.Code)
	assert.Equal(t, "success", response.Meta.Status)
	assert.Equal(t, string(data), response.Data)
}

func TestSend(t *testing.T) {
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	user := User{ID: 1, Name: "John"}
	data, err := json.Marshal(user)
	assert.Nil(t, err)

	response := ResponseJSON("success", 200, "success", string(data))
	assert.Equal(t, "success", response.Meta.Message)
	assert.Equal(t, 200, response.Meta.Code)
	assert.Equal(t, "success", response.Meta.Status)
	assert.Equal(t, string(data), response.Data)

	response.Send(nil)
}

func TestValidationError(t *testing.T) {
	type User struct {
		Name string `json:"name" validate:"required"`
	}

	data := `{"name":"dev"}`

	var l User
	json.Unmarshal([]byte(data), &l)
	err := ValidationError(l)
	assert.NoError(t, err, "should not have error")
}
