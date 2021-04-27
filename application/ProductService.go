package application

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
	ProductFactory    entity.ProductFactory
}

func NewProductService(ProductRepository repository.ProductRepository,
	ProductFactory entity.ProductFactory) *ProductService {
	return &ProductService{ProductRepository: ProductRepository,
		ProductFactory: ProductFactory}
}

func (s *ProductService) GetByID(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "ProductId invalid."))
		return
	}

	var product, err = s.ProductRepository.FindByID(id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "Product not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, product)
	} else {
		c.JSON(200, s.ProductFactory.Create(product, fields))
	}
}

func (s *ProductService) GetAll(c *gin.Context) {
	var products, err = s.ProductRepository.FindAll()
	if err != nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "Products not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, products)
	} else {
		c.JSON(200, s.ProductFactory.CreateList(products, fields))
	}

}
