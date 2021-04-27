package application

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-api/domain/service"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type OrderRequestResource struct {
	TableID   string `json:"tableID"`
	ProductID string `json:"productID"`
	Count     string `json:"count"`
	Comments  string `json:"comments"`
	ID        string `json:"id"`
}

type OrderService struct {
	OrderService    *service.OrderService
	OrderRepository repository.OrderRepository
}

func NewOrderService(
	orderService *service.OrderService,
	orderRepository repository.OrderRepository) *OrderService {
	return &OrderService{
		OrderService:    orderService,
		OrderRepository: orderRepository}
}

func (s *OrderService) GetOrderDetails(c *gin.Context) {
	orderNumberIDStr := c.Query("orderNumberID")
	if orderNumberIDStr != "" {
		orderNumberIDInt, paramErr := strconv.ParseInt(orderNumberIDStr, 0, 64)
		if paramErr != nil {
			log.Print(paramErr)
			c.JSON(http.StatusBadRequest, exception.CreateError(exception.CodeValueInvalid, "orderNumberID is invalid."))
			return
		}
		var orders, err = s.OrderRepository.FindDetailByOrderNumberID(orderNumberIDInt)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusNotFound, exception.CreateError(exception.CodeNotFound, "order not found."))
			return
		}
		var fields = c.Query("fields")
		if fields == "" {
			c.JSON(http.StatusOK, orders.ToArray())
		} else {
			c.JSON(http.StatusOK, orders.ToSlice(fields))
		}
		return
	}
	c.JSON(http.StatusBadRequest, exception.CreateError(exception.CodeValueInvalid, "orderNumberID is required."))
}

func (s *OrderService) PutOrders(c *gin.Context) {
	orderRequest := entity.OrderRequest{}
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, exception.GetError(exception.CodeValueInvalid))
		return
	}

	orderNumberID, err := s.OrderService.Orders(orderRequest)
	if err != nil {
		switch err.ErrorCode {
		case exception.CodeValueInvalid:
			c.JSON(http.StatusBadRequest, err)
		default:
			c.JSON(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"orderNumberID": orderNumberID})
}
