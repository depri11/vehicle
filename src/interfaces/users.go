package interfaces

import (
	"net/http"

	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/helper"
)

type UserRepository interface {
	FindAll() (*models.Users, error)
	Save(user *models.User) (*models.User, error)
	GetUserID(ID int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(ID uint) error
	Search(search string) (*models.Users, error)
	Sort(sort string) (*models.Users, error)
}

type UserService interface {
	FindAll() (*helper.Res, error)
	FindByEmail(email string) (*helper.Res, error)
	FindByID(id int) (*helper.Res, error)
	RegisterUser(user *models.User) (*helper.Res, error)
	UpdateUser(id int, user *models.User, r *http.Request) (*helper.Res, error)
	Delete(id int, r *http.Request) (*helper.Res, error)
	Sort(sort string) (*helper.Res, error)
	Search(search string) (*helper.Res, error)
}
