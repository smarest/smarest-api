package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type AreaRepository interface {
	FindByID(id int64) (*entity.Area, error)
	FindByRestaurantID(restId int64) ([]entity.Area, error)
	FindAll() ([]entity.Area, error)
}
