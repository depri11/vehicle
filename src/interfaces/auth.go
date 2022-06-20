package interfaces

import (
	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/helper"
)

type AuthService interface {
	Login(user models.User) *helper.Res
}
