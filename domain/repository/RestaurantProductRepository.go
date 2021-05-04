package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type RestaurantProductRepository interface {
	FindAvailableProductsByCategoryID(resId int64, categoryID int64) (*entity.ProductList, error)
	FindAvailableProductByID(resId int64, productID int64) (*entity.Product, error)
	FindAvailableProductsByRestaurantIDAndIDs(resId int64, productIDs []int64) (*entity.ProductList, error)
}
