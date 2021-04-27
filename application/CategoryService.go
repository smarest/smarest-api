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

type CategoryService struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(CategoryRepository repository.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryRepository: CategoryRepository}
}

func (s *CategoryService) GetByID(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "category invalid."))
		return
	}

	var result, err = s.CategoryRepository.FindByID(id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "category not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, result)
	} else {
		c.JSON(200, result.ToSlice(fields))
	}
}

func (s *CategoryService) GetAll(c *gin.Context) {
	cateType := c.Query("type")
	available := c.Query("available")
	var results entity.CategoryList
	var err error
	if cateType == "" {
		results, err = s.CategoryRepository.FindAll()
	} else {
		results, err = s.CategoryRepository.FindByType(cateType)
	}

	if available == "true" {
		results = results.GetAvailable()
	}

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusNotFound, exception.CreateError(exception.CodeNotFound, "category not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, results.ToArray())
	} else {
		c.JSON(200, results.ToSlice(fields))
	}
}
