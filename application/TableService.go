package application

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type TableService struct {
	TableRepository repository.TableRepository
	TableFactory    entity.TableFactory
}

func NewTableService(TableRepository repository.TableRepository,
	TableFactory entity.TableFactory) *TableService {
	return &TableService{
		TableRepository: TableRepository,
		TableFactory:    TableFactory}
}

func (s *TableService) GetByID(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "tableID invalid."))
		return
	}

	var result, err = s.TableRepository.FindByID(id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "Table not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, result)
	} else {
		c.JSON(200, s.TableFactory.Create(result, fields))
	}
}

func (s *TableService) GetAll(c *gin.Context) {
	var results, err = s.TableRepository.FindAll()
	if err != nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "Tables not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, results)
	} else {
		c.JSON(200, s.TableFactory.CreateList(results, fields))
	}
}
