package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type RestaurantRepository interface {
	FindByID(id int64) (*entity.Restaurant, error)
	FindAll() ([]entity.Restaurant, error)
	FindProductByID(restaurantID int64, productID int64) (*entity.Product, error)
	FindProductByIDs(restaurantID int64, productIDs []int64) (entity.ProductList, error)
}

type RestaurantLoginRepository interface {
	FindByID(id int64) (*entity.Restaurant, error)
	FindByIDAndAccessKey(id int64, accessKey string) (*entity.Restaurant, error)
}
