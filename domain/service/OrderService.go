package service

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	common_entity "github.com/smarest/smarest-common/domain/entity"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type OrderService struct {
	AreaRepository              repository.AreaRepository
	OrderRepository             repository.OrderRepository
	RestaurantProductRepository repository.RestaurantProductRepository
	ProductRepository           repository.ProductRepository
	TableRepository             repository.TableRepository
	OrderFactory                entity.OrderFactory
}

func NewOrderService(
	areaRepository repository.AreaRepository,
	orderRepository repository.OrderRepository,
	productRepository repository.ProductRepository,
	restaurantProductRepository repository.RestaurantProductRepository,
	tableRepository repository.TableRepository,
	orderFactory entity.OrderFactory) *OrderService {
	return &OrderService{
		AreaRepository:              areaRepository,
		OrderRepository:             orderRepository,
		ProductRepository:           productRepository,
		RestaurantProductRepository: restaurantProductRepository,
		TableRepository:             tableRepository,
		OrderFactory:                orderFactory,
	}
}

func (s *OrderService) GetRestaurantOrdersByAreaIDAndGroupByOrderNumberID(restaurantID int64, areaID int64) ([]entity.OrderGroupByOrderNumberID, *exception.Error) {
	area, aErr := s.AreaRepository.FindAvailableByID(areaID)
	if aErr != nil {
		if aErr == sql.ErrNoRows {
			return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("AreaID=[%d] not found. ", areaID))
		}
		return nil, exception.CreateErrorWithRootCause(exception.CodeSystemError, fmt.Sprintf("Can not get AreaID=[%d]", areaID), aErr)
	}

	if area.RestaurantID != restaurantID {
		return nil, exception.CreateError(exception.CodeValueInvalid, fmt.Sprintf("AreaID=[%d] not belong to RestaurantID=[%d]. ", areaID, restaurantID))
	}

	var orders, err = s.OrderRepository.FindByAreaIDAndGroupByOrderNumberID(areaID)
	if err != nil {
		return nil, exception.CreateErrorWithRootCause(exception.CodeSystemError, "Can not get Orders. ", err)
	}

	return orders, nil
}

func (s *OrderService) GetOrderDetailsByRestaurantIDAndOrderNumberID(restaurantID int64, orderNumberID int64) (*entity.OrderDetailList, *exception.Error) {
	orderNumber, orderNumberErr := s.OrderRepository.FindOrderNumber(orderNumberID)
	if orderNumberErr != nil {
		if orderNumberErr == sql.ErrNoRows {
			return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("OrderNumberID=[%d] not found.", orderNumberID))
		}
		return nil, exception.CreateErrorWithRootCause(exception.CodeSystemError, fmt.Sprintf("Can not get OrderNumberID=[%d]", orderNumberID), orderNumberErr)
	}

	// order number is not belong to restaurant
	if orderNumber.RestaurantID != restaurantID {
		return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("OrderNumberID=[%d] not belong to RestaurantID=[%d].", orderNumberID, restaurantID))
	}

	var orders, err = s.OrderRepository.FindByOrderNumberID(orderNumberID)
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}

	// get Products and Tables
	productIDs, tableIds := orders.GetDistinctProductIDsAndTableIDs()
	// get Products
	productList, err := s.ProductRepository.FindByIDs(productIDs)
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}
	// get Table
	tableList, err := s.TableRepository.FindByIDs(tableIds)
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}

	orderDetails := entity.CreateEmptyOrderDetailList()
	for _, order := range orders.Orders {
		product := productList.FindByID(order.ProductID)
		if product == nil {
			return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("Product not found. productID=[%d]", order.ProductID))
		}

		table := tableList.FindByID(order.TableID)
		if table == nil {
			return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("Table not found. tableID=[%d]", order.TableID))
		}
		orderDetails.Add(entity.NewOrderDetail(&order, table, product))
	}

	return orderDetails, nil
}

func (s *OrderService) GetOrderDetailsByRestaurantID(restaurantID int64) (*entity.OrderDetailList, *exception.Error) {
	var orders, err = s.OrderRepository.FindByRestaurantID(restaurantID)
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}

	// get Products and Tables
	productIDs, tableIds := orders.GetDistinctProductIDsAndTableIDs()
	// get Products
	productList, err := s.ProductRepository.FindByIDs(productIDs)
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}
	// get Table
	tableList, err := s.TableRepository.FindByIDs(tableIds)
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}

	orderDetails := entity.CreateEmptyOrderDetailList()
	for _, order := range orders.Orders {
		product := productList.FindByID(order.ProductID)
		if product == nil {
			return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("Product not found. productID=[%d]", order.ProductID))
		}

		table := tableList.FindByID(order.TableID)
		if table == nil {
			return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("Table not found. tableID=[%d]", order.TableID))
		}
		orderDetails.Add(entity.NewOrderDetail(&order, table, product))
	}

	return orderDetails, nil
}

func (s *OrderService) Orders(orderRequest common_entity.OrderRequest) (*entity.OrderNumber, *exception.Error) {
	if orderRequest.WaiterID == nil || orderRequest.RestaurantID == nil {
		return nil, exception.CreateError(exception.CodeValueInvalid, "orderNumberID, restaurantID, waiterID is required.")
	}

	if orderRequest.OrderNumberID == nil && orderRequest.IsEmpty() {
		return nil, exception.CreateError(exception.CodeValueInvalid, "order Data is required.")
	}

	// get Products and Tables
	productIDs, tableIds := orderRequest.GetDistinctProductIDsAndTableIDs()
	// get Products
	productList, err := s.RestaurantProductRepository.FindAvailableProductsByRestaurantIDAndIDs(*orderRequest.RestaurantID, productIDs)
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}
	// get Table
	tableList, err := s.TableRepository.FindAvailableByRestaurantIDAndIDs(*orderRequest.RestaurantID, tableIds)
	if err != nil {
		return nil, exception.CreateError(exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
	}

	// 1. is new order
	if orderRequest.OrderNumberID == nil {
		// get order number
		orderNumber, err := s.OrderRepository.RegisterOrderNumber(*orderRequest.RestaurantID)
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
				return nil, s.RollbackAndCreateError(orderNumber, exception.CodeValueInvalid, fmt.Sprintf("Product not found. ProductID=%d", orderRecord.ProductID))
			}

			table := tableList.FindByID(orderRecord.TableID)
			if table == nil {
				return nil, s.RollbackAndCreateError(orderNumber, exception.CodeValueInvalid, fmt.Sprintf("Table not found. TableID=%d", orderRecord.TableID))
			}

			_, err = s.OrderRepository.RegisterOrder(s.OrderFactory.CreateForRegister(orderNumber.ID, *orderRequest.WaiterID, orderRecord, product))
			if err != nil {
				fmt.Print(err.Error())
				return nil, s.RollbackAndCreateError(orderNumber, exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
			}
		}

		s.OrderRepository.Commit()
		if err != nil {
			return nil, s.RollbackAndCreateError(orderNumber, exception.CodeSystemError, fmt.Sprintf("System error: [%s]", err))
		}
		// end transaction
		return orderNumber, nil
	}

	// 2. check orderNumberID
	orderNumber, err := s.OrderRepository.FindOrderNumber(*orderRequest.OrderNumberID)
	if err != nil {
		return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("System error: [%s]", err))
	}
	// order number is not found
	if orderNumber == nil {
		return nil, exception.CreateError(exception.CodeValueInvalid, fmt.Sprintf("OrderNumberID=[%d] not found. ", *orderRequest.OrderNumberID))
	}
	// order number not belong to restaurant
	if orderNumber.RestaurantID != *orderRequest.RestaurantID {
		return nil, exception.CreateError(exception.CodeNotFound, fmt.Sprintf("OrderNumberID=[%d] not belong to RestaurantID=[%d].", orderNumber.ID, *orderRequest.RestaurantID))
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

		table := tableList.FindByID(updateOrder.TableID)
		if table == nil {
			return nil, s.RollbackAndCreateError(nil, exception.CodeValueInvalid, fmt.Sprintf("Table not found. TableID=%d", updateOrder.TableID))
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
	log.Printf("len: %d", len(willBeUpdateOrders)+len(willBeRegisterOrders))
	if len(willBeUpdateOrders)+len(willBeRegisterOrders) == 0 {
		_, errs := s.OrderRepository.DeleteOrderNumber(orderNumber)
		log.Printf("%s", errs)
	}
	return orderNumber, nil
}

func (s *OrderService) RollbackAndCreateError(orderNumber *entity.OrderNumber, code int, message string) *exception.Error {
	log.Print(message)
	s.OrderRepository.Rollback()
	if orderNumber != nil {
		s.OrderRepository.DeleteOrderNumber(orderNumber)
	}
	return exception.CreateError(code, message)
}
