package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type OrderRepository interface {
	DAO
	FindByAreaIDAndGroupByOrderNumberID(id int64) ([]entity.OrderGroupByOrderNumberID, error)
	FindDetailByOrderNumberID(orderNumberID int64) (entity.OrderDetailList, error)
	FindByOrderNumberID(orderNumberID int64) (entity.OrderList, error)
	RegisterOrder(order entity.Order) (int64, error)
	UpdateOrder(order entity.Order) (int64, error)
	RegisterOrderNumber() (int64, error)
	DeleteOrderNumber(orderNumberID int64) (int64, error)
	FindOrderNumber(orderNumberID int64) (*entity.OrderNumber, error)
	DeleteByOrderNumberIDAndIDNotIn(orderNumberID int64, ids []int64) (int64, error)
}
