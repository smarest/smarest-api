package application

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type RestaurantService struct {
	RestaurantRepository repository.RestaurantRepository
	ProductRepository    repository.ProductRepository
	AreaRepository       repository.AreaRepository
	RestaurantFactory    entity.RestaurantFactory
	ProductFactory       entity.ProductFactory
	AreaFactory          entity.AreaFactory
}

func NewRestaurantService(
	restaurantRepository repository.RestaurantRepository,
	productRepository repository.ProductRepository,
	areaRepository repository.AreaRepository,
	restaurantFactory entity.RestaurantFactory,
	productFactory entity.ProductFactory,
	areaFactory entity.AreaFactory) *RestaurantService {
	return &RestaurantService{RestaurantRepository: restaurantRepository,
		ProductRepository: productRepository,
		AreaRepository:    areaRepository,
		RestaurantFactory: restaurantFactory,
		ProductFactory:    productFactory,
		AreaFactory:       areaFactory,
	}
}

func (s *RestaurantService) GetByID(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "restaurantId invalid."))
		return
	}

	var restaurant, err = s.RestaurantRepository.FindByID(id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "restaurant not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, restaurant)
	} else {
		c.JSON(200, s.RestaurantFactory.Create(restaurant, fields))
	}
}

func (s *RestaurantService) GetAll(c *gin.Context) {
	var restaurants, err = s.RestaurantRepository.FindAll()
	if err != nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "restaurant not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, restaurants)
	} else {
		c.JSON(200, s.RestaurantFactory.CreateList(restaurants, fields))
	}
}

func (s *RestaurantService) GetProducts(c *gin.Context) {
	var results []entity.Product
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "restaurantId invalid."))
		return
	}

	var products, err = s.ProductRepository.FindByRestaurantID(id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "product not found."))
		return
	}

	categoryIDStr := c.Query("categoryID")
	if categoryIDStr != "" {
		results = []entity.Product{}
		for _, item := range products {
			if fmt.Sprint(item.CategoryID) == categoryIDStr {
				results = append(results, item)
			}
		}
	} else {
		results = products
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, results)
	} else {
		c.JSON(200, s.ProductFactory.CreateList(results, fields))
	}
}

func (s *RestaurantService) GetAreas(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "restaurantId invalid."))
		return
	}

	var areas, err = s.AreaRepository.FindByRestaurantID(id)
	if err != nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "area not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, areas)
	} else {
		c.JSON(200, s.AreaFactory.CreateList(areas, fields))
	}
}
