package application

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smarest/smarest-common/domain/entity"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type PortalService struct {
	Bean *Bean
}

type Response interface {
	ToSlice() interface{}
	ToSliceWithFields(fields string) interface{}
}

type Responses interface {
	ToSlice() []interface{}
	ToSliceWithFields(fields string) []interface{}
}

func NewPortalService(bean *Bean) *PortalService {
	return &PortalService{bean}
}

func (s *PortalService) GetRestaurantsByRestaurantGroupID(c *gin.Context) {
	groupID, paramErr := strconv.ParseInt(c.Params.ByName("groupID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateErrorWithRootCause(exception.CodeValueInvalid, "groupID invalid.", paramErr))
		return
	}

	var restaurants, err = s.Bean.RestaurantService.FindAvailableByRestaurantGroupID(groupID)
	if err != nil {
		s.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, restaurants.ToSlice(c.Query("fields")))
}

func (s *PortalService) GetRestaurantByCode(c *gin.Context) {
	var restaurant, err = s.Bean.RestaurantService.FindAvailableByCode(c.Query("code"))
	if err != nil {
		s.HandlerError(c, err)
		return
	}
	c.JSON(http.StatusOK, restaurant.ToSlice(c.Query("fields")))
}

func (s *PortalService) GetRestaurantAreas(c *gin.Context) {
	restaurantIDInt, paramErr := strconv.ParseInt(c.Params.ByName("restaurantID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateErrorWithRootCause(exception.CodeValueInvalid, "restaurantID invalid.", paramErr))
		return
	}

	var areas, err = s.Bean.AreaService.FindAvailableByRestaurantID(restaurantIDInt)
	if err != nil {
		s.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, areas.ToSlice(c.Query("fields")))
}

func (s *PortalService) GetRestaurantTablesByAreaID(c *gin.Context) {
	restaurantID, paramErr := strconv.ParseInt(c.Params.ByName("restaurantID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateErrorWithRootCause(exception.CodeValueInvalid, "restaurantID invalid.", paramErr))
		return
	}
	areaID, paramErr := strconv.ParseInt(c.Params.ByName("areaID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateErrorWithRootCause(exception.CodeValueInvalid, "areaID invalid.", paramErr))
		return
	}

	var tables, err = s.Bean.TableService.FindAvailableByRestaurantIDAndAreaID(restaurantID, areaID)
	if err != nil {
		s.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, tables.ToSlice(c.Query("fields")))
}

func (s *PortalService) GetRestaurantGroupCategories(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("groupID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateErrorWithRootCause(exception.CodeValueInvalid, "groupID invalid.", paramErr))
		return
	}

	var categories, err = s.Bean.CategoryService.FindAvailableByRestaurantGroupIDAndType(id, c.Query("type"))
	if err != nil {
		s.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, categories.ToSlice(c.Query("fields")))
}

func (s *PortalService) GetRestaurantProductsByCategoryID(c *gin.Context) {
	restID, paramErr := strconv.ParseInt(c.Params.ByName("restaurantID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateErrorWithRootCause(exception.CodeValueInvalid, "restaurantID invalid.", paramErr))
		return
	}

	categoryID, paramErr := strconv.ParseInt(c.Params.ByName("categoryID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateErrorWithRootCause(exception.CodeValueInvalid, "categoryID invalid.", paramErr))
		return
	}

	var products, err = s.Bean.RestaurantService.GetAvailableProductsByIDAndCategoryID(restID, categoryID)
	if err != nil {
		s.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, products.ToSlice(c.Query("fields")))
}

func (s *PortalService) GetRestaurantGroupCommentsByProductID(c *gin.Context) {
	groupID, paramErr := strconv.ParseInt(c.Params.ByName("groupID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateErrorWithRootCause(exception.CodeValueInvalid, "groupID invalid.", paramErr))
		return
	}
	productID, paramErr := strconv.ParseInt(c.Params.ByName("productID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateErrorWithRootCause(exception.CodeValueInvalid, "productID invalid.", paramErr))
		return
	}

	var comments, err = s.Bean.CommentService.FindAvailableByRestaurantGroupIDAndProductID(groupID, productID)
	if err != nil {
		s.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, comments.ToSlice(c.Query("fields")))
}

func (s *PortalService) GetRestaurantOrdersByAreaIDAndGroupByOrderNumberID(c *gin.Context) {
	restID, paramErr := strconv.ParseInt(c.Params.ByName("restaurantID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateError(exception.CodeValueInvalid, "restaurantID invalid."))
		return
	}

	areaId, paramErr := strconv.ParseInt(c.Params.ByName("areaID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateError(exception.CodeValueInvalid, "areaID invalid."))
		return
	}

	var orders, err = s.Bean.OrderService.GetRestaurantOrdersByAreaIDAndGroupByOrderNumberID(restID, areaId)
	if err != nil {
		s.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (s *PortalService) GetRestaurantOrderDetailsByOrderNumberID(c *gin.Context) {
	restID, paramErr := strconv.ParseInt(c.Params.ByName("restaurantID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateError(exception.CodeValueInvalid, "restaurantID invalid."))
		return
	}
	orderNumberID, paramErr := strconv.ParseInt(c.Query("orderNumberID"), 0, 64)
	if paramErr != nil {
		s.HandlerError(c, exception.CreateError(exception.CodeValueInvalid, "orderNumberID invalid."))
		return
	}

	var orders, err = s.Bean.OrderService.GetOrderDetailsByRestaurantIDAndOrderNumberID(restID, orderNumberID)
	if err != nil {
		s.HandlerError(c, err)
		return
	}

	c.JSON(http.StatusOK, orders.ToSlice(c.Query("fields")))
}

func (s *PortalService) PutRestaurantOrders(c *gin.Context) {
	orderRequest := entity.OrderRequest{}
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		s.HandlerError(c, exception.GetError(exception.CodeValueInvalid))
		return
	}

	orderNumber, err := s.Bean.OrderService.Orders(orderRequest)
	if err != nil {
		s.HandlerError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"orderNumberID": orderNumber.ID})
}

func (s *PortalService) HandlerError(c *gin.Context, err *exception.Error) {
	log.Printf("Message=[%s], error=[%s]", err.ErrorMessage, err.RootCause)
	switch err.ErrorCode {
	case exception.CodeNotFound:
		c.JSON(http.StatusNotFound, err)
	case exception.CodeValueInvalid:
		c.JSON(http.StatusBadRequest, err)
	default:
		c.JSON(http.StatusInternalServerError, err)
	}
}
