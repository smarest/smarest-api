package entity

import (
	"github.com/smarest/smarest-common/domain/entity"
	"github.com/smarest/smarest-common/domain/value"
)

type OrderFactory struct{}

func NewOrderFactory() OrderFactory {
	return OrderFactory{}
}

func (f *OrderFactory) CreateForRegister(orderNumberID int64, waiterID string, orderRequestRecord entity.OrderRequestRecord, product *Product) Order {
	return f.createFrom(-1, orderNumberID, waiterID, orderRequestRecord, product)
}

func (f *OrderFactory) CreateForUpdate(orderNumberID int64, waiterID string, orderRequestRecord entity.OrderRequestRecord, product *Product) Order {
	return f.createFrom(*orderRequestRecord.ID, orderNumberID, waiterID, orderRequestRecord, product)
}

func (f *OrderFactory) createFrom(orderID int64, orderNumberID int64, waiterID string, orderRequestRecord entity.OrderRequestRecord, product *Product) Order {
	return Order{
		ID:            orderID,
		WaiterID:      waiterID,
		ChefID:        nil,
		OrderNumberID: orderNumberID,
		TableID:       orderRequestRecord.TableID,
		ProductID:     product.ID,
		Count:         product.QuantityOnSingleOrder, //TODO fix it to manual
		Comments:      orderRequestRecord.Comments,
		Status:        product.DefaultStatus,
		Price:         product.Price,
		OrderTime:     value.NewDateTime(),
	}
}
