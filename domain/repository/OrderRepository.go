package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type OrderRepository interface {
	DAO
	FindByAreaIDAndGroupByOrderNumberID(id int64) ([]entity.OrderGroupByOrderNumberID, error)
	FindByOrderNumberID(orderNumberID int64) (*entity.OrderList, error)
	FindByRestaurantID(restaurantID int64) (*entity.OrderList, error)
	RegisterOrder(order entity.Order) (int64, error)
	UpdateOrder(order entity.Order) (int64, error)
	DeleteByOrderNumberIDAndIDNotIn(orderNumberID int64, ids []int64) (int64, error)
	RegisterOrderNumber(restaurantID int64) (*entity.OrderNumber, error)
	DeleteOrderNumber(orderNumber *entity.OrderNumber) (int64, error)
	FindOrderNumber(orderNumberID int64) (*entity.OrderNumber, error)
}
