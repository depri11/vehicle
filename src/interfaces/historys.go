package interfaces

import (
	"net/http"

	"github.com/depri11/vehicle/src/database/models"
	"github.com/depri11/vehicle/src/helper"
)

type HistoryRepository interface {
	FindAll() (*models.Historyss, error)
	GetID(ID int) (*models.Historys, error)
	Save(history *models.Historys) (*models.Historys, error)
	Update(history *models.Historys) (*models.Historys, error)
	Delete(ID int) error
	Sort(sort string) (*models.Historyss, error)
	Search(search string) (*models.Historyss, error)
}

type HistoryService interface {
	FindAll() (*helper.Res, error)
	FindByID(id int) (*helper.Res, error)
	Create(user *models.Historys, r *http.Request) (*helper.Res, error)
	Update(id int, vehicle *models.Historys, r *http.Request) (*helper.Res, error)
	Delete(id int, r *http.Request) (*helper.Res, error)
	Sort(sort string) (*helper.Res, error)
	Search(search string) (*helper.Res, error)
}
