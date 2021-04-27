package service

import (
	"fmt"
	"log"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type OrderService struct {
	OrderRepository      repository.OrderRepository
	RestaurantRepository repository.RestaurantRepository
	TableRepository      repository.TableRepository
	OrderFactory         entity.OrderFactory
}

func NewOrderService(
	orderRepository repository.OrderRepository,
	orderFactory entity.OrderFactory,
	restaurantRepository repository.RestaurantRepository,
	tableRepository repository.TableRepository) *OrderService {
	return &OrderService{
		OrderRepository:      orderRepository,
		OrderFactory:         orderFactory,
		RestaurantRepository: restaurantRepository,
		TableRepository:      tableRepository}
}

func (s *OrderService) Orders(orderRequest entity.OrderRequest) (*int64, *exception.Error) {
	if orderRequest.WaiterID == nil || orderRequest.RestaurantID == nil || (len(orderRequest.OrderRequestRecords) < 1 && orderRequest.OrderNumberID == nil) {
		return nil, exception.CreateError(exception.CodeValueInvalid, "orderNumberID, restaurantID, waiterID is required.")
	}

	// get Product
	productList, err := s.RestaurantRepository.FindProductByIDs(*orderRequest.RestaurantID, orderRequest.GetDistinctProductIDs())
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}

	// 1. is new order
	if orderRequest.OrderNumberID == nil {
		// get order number
		orderNumberID, err := s.OrderRepository.RegisterOrderNumber()
		if err != nil {
			return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
		}

		// begin transaction
		err = s.OrderRepository.Begin()
		if err != nil {
			return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
		}

		for _, orderRecord := range orderRequest.OrderRequestRecords {
			product := productList.FindByID(orderRecord.ProductID)
			if product == nil {
				return nil, s.RollbackAndCreateError(&orderNumberID, exception.CodeValueInvalid, fmt.Sprintf("Product not found. ProductID=%d", orderRecord.ProductID))
			}

			_, err = s.OrderRepository.RegisterOrder(s.OrderFactory.CreateForRegister(orderNumberID, *orderRequest.WaiterID, orderRecord, product))
			if err != nil {
				fmt.Print(err.Error())
				return nil, s.RollbackAndCreateError(&orderNumberID, exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
			}
		}

		s.OrderRepository.Commit()
		if err != nil {
			return nil, s.RollbackAndCreateError(&orderNumberID, exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
		}
		// end transaction
		return &orderNumberID, nil
	}

	// 2. check orderNumberID
	orderNumber, err := s.OrderRepository.FindOrderNumber(*orderRequest.OrderNumberID)
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}
	// order number is not found
	if orderNumber == nil {
		return nil, exception.CreateError(exception.CodeValueInvalid, fmt.Sprintf("OrderNumberID=%d not found. ", *orderRequest.OrderNumberID))
	}

	// 3. is edit order or delete order
	willBeRegisterOrders, willBeUpdateOrders, willBeUpdateOrderIDs := orderRequest.Discover()

	// begin transaction
	err = s.OrderRepository.Begin()
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}

	// 3.1
	// update
	for _, updateOrder := range willBeUpdateOrders {
		product := productList.FindByID(updateOrder.ProductID)
		if product == nil {
			return nil, s.RollbackAndCreateError(nil, exception.CodeValueInvalid, fmt.Sprintf("Product not found. ProductID=%d", updateOrder.ProductID))
		}
		_, err = s.OrderRepository.UpdateOrder(s.OrderFactory.CreateForUpdate(orderNumber.ID, *orderRequest.WaiterID, updateOrder, product))
		if err != nil {
			return nil, s.RollbackAndCreateError(nil, exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err.Error()))
		}
	}

	// 3.2 delete not use
	// TODO log or memo to database
	_, err = s.OrderRepository.DeleteByOrderNumberIDAndIDNotIn(orderNumber.ID, willBeUpdateOrderIDs)
	if err != nil {
		return nil, s.RollbackAndCreateError(nil, exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}

	// 3.3
	// register
	for _, registerOrder := range willBeRegisterOrders {
		product := productList.FindByID(registerOrder.ProductID)
		if product == nil {
			fmt.Print(err.Error())
			return nil, s.RollbackAndCreateError(nil, exception.CodeValueInvalid, fmt.Sprintf("Product not found. ProductID=%d", registerOrder.ProductID))
		}
		_, err = s.OrderRepository.RegisterOrder(s.OrderFactory.CreateForRegister(orderNumber.ID, *orderRequest.WaiterID, registerOrder, product))
		if err != nil {
			return nil, s.RollbackAndCreateError(nil, exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
		}
	}

	// end transaction
	s.OrderRepository.Commit()
	if err != nil {
		return nil, s.RollbackAndCreateError(nil, exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}

	// 3.4 if update list and register list size = 0 this request meaning delete order
	if len(willBeUpdateOrderIDs)+len(willBeRegisterOrders) == 0 {
		s.OrderRepository.DeleteOrderNumber(orderNumber.ID)
	}
	return &orderNumber.ID, nil
}

func (s *OrderService) RollbackAndCreateError(orderNumberID *int64, code int, message string) *exception.Error {
	log.Print(message)
	s.OrderRepository.Rollback()
	if orderNumberID != nil {
		s.OrderRepository.DeleteOrderNumber(*orderNumberID)
	}
	return exception.CreateError(code, message)
}
