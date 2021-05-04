package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type AreaRepository interface {
	FindByID(id int64) (*entity.Area, error)
	FindAvailableByID(id int64) (*entity.Area, error)
	FindAvailableByRestaurantID(restId int64) (*entity.AreaList, error)
}
