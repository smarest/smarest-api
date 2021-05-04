package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type RestaurantRepository interface {
	FindByID(id int64) (*entity.Restaurant, error)
	FindByRestaurantGroupID(restaurantGroupID int64) (*entity.RestaurantList, error)
	FindAvailableByRestaurantGroupID(restaurantGroupID int64) (*entity.RestaurantList, error)
	FindByCode(code string) (*entity.Restaurant, error)
	FindAll() (*entity.RestaurantList, error)
}
