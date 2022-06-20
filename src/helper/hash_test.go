package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "12345"
	result, _ := HashPassword(password)
	assert.NotEqual(t, password, result, "Expected hash to be created")
}

func TestCheckPassword(t *testing.T) {
	password := "12345"
	hash, _ := HashPassword(password)
	result := CheckPassword(hash, password)
	assert.Equal(t, true, result, "Expected password to be correct")
}
