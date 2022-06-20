package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	result := NewToken(1, "dev@gmail.com", "admin")
	assert.Equal(t, 1, result.Id)
}

func TestCreate(t *testing.T) {
	example := &claims{
		Id:    1,
		Email: "dev@gmail.com",
		Role:  "admin",
	}

	result, _ := example.Create()
	assert.NotEqual(t, nil, result, "Expected token to be created")
}

func TestCheckToken(t *testing.T) {
	example := &claims{
		Id:    1,
		Email: "dev@gmail.com",
		Role:  "admin",
	}

	result, _ := example.Create()
	result2, _ := CheckToken(result)
	assert.Equal(t, example.Id, result2.Id, "Expected id to be correct")
	assert.Equal(t, example.Email, result2.Email, "Expected email to be correct")
	assert.Equal(t, example.Role, result2.Role, "Expected role to be correct")
}
