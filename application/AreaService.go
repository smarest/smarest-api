package application

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type AreaService struct {
	AreaRepository  repository.AreaRepository
	AreaFactory     entity.AreaFactory
	TableRepository repository.TableRepository
	TableFactory    entity.TableFactory
	OrderRepository repository.OrderRepository
}

func NewAreaService(AreaRepository repository.AreaRepository,
	AreaFactory entity.AreaFactory, TableRepository repository.TableRepository,
	TableFactory entity.TableFactory, OrderRepository repository.OrderRepository) *AreaService {
	return &AreaService{
		AreaRepository:  AreaRepository,
		AreaFactory:     AreaFactory,
		TableRepository: TableRepository,
		TableFactory:    TableFactory,
		OrderRepository: OrderRepository}
}

func (s *AreaService) GetByID(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "AreaID invalid."))
		return
	}

	var result, err = s.AreaRepository.FindByID(id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "Area not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, result)
	} else {
		c.JSON(200, s.AreaFactory.Create(result, fields))
	}
}

func (s *AreaService) GetAll(c *gin.Context) {
	var results, err = s.AreaRepository.FindAll()
	if err != nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "Areas not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, results)
	} else {
		c.JSON(200, s.AreaFactory.CreateList(results, fields))
	}
}

func (s *AreaService) GetTables(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "areaID invalid."))
		return
	}

	var tables, err = s.TableRepository.FindByAreaID(id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "tables not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, tables)
	} else {
		c.JSON(200, s.TableFactory.CreateList(tables, fields))
	}
}

func (s *AreaService) GetOrders(c *gin.Context) {
	areaId, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "areaID invalid."))
		return
	}

	var orders, err = s.OrderRepository.FindByAreaIDAndGroupByOrderNumberID(areaId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusNotFound, exception.CreateError(exception.CodeNotFound, "order not found."))
		return
	}

	c.JSON(http.StatusOK, orders)
}
