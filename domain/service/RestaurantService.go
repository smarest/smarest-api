package service

import (
	"fmt"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type RestaurantService struct {
	RestaurantRepository        repository.RestaurantRepository
	RestaurantProductRepository repository.RestaurantProductRepository
}

func NewRestaurantService(restaurantRepository repository.RestaurantRepository, restaurantProductRepository repository.RestaurantProductRepository) *RestaurantService {
	return &RestaurantService{restaurantRepository, restaurantProductRepository}
}

func (s *RestaurantService) FindAvailableByCode(code string) (*entity.Restaurant, *exception.Error) {
	if code == "" {
		return nil, exception.CreateError(exception.CodeValueInvalid, "RestaurantCode is required")
	}
	var result, err = s.RestaurantRepository.FindByCode(code)
	if err != nil {
		return nil, exception.CreateErrorWithRootCause(exception.CodeNotFound, fmt.Sprintf("Restaurant not found. code=[%s]", code), err)
	}
	return result, nil
}

func (s *RestaurantService) FindAvailableByRestaurantGroupID(groupID int64) (*entity.RestaurantList, *exception.Error) {
	var results, err = s.RestaurantRepository.FindAvailableByRestaurantGroupID(groupID)
	if err != nil {
		return nil, exception.CreateErrorWithRootCause(exception.CodeNotFound, fmt.Sprintf("Restaurants not found. restaurantGroupID=[%d]", groupID), err)
	}

	return results, nil
}

func (s *RestaurantService) GetAvailableProductsByIDAndCategoryID(restaurantID int64, categoryID int64) (*entity.ProductList, *exception.Error) {
	var results, err = s.RestaurantProductRepository.FindAvailableProductsByCategoryID(restaurantID, categoryID)
	if err != nil {
		return nil, exception.CreateErrorWithRootCause(exception.CodeNotFound, fmt.Sprintf("Restaurants not found. restaurantID=[%d], categoryID=[%d]", restaurantID, categoryID), err)
	}

	return results, nil
}
