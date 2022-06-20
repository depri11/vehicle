package helper

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
