package service

import (
	"fmt"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
	"github.com/smarest/smarest-common/util"
)

type CategoryService struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository}
}

func (s *CategoryService) FindAvailableByRestaurantGroupIDAndType(restaurantGroupID int64, cateType string) (*entity.CategoryList, *exception.Error) {
	var results entity.CategoryList
	var err error

	switch cateType {
	case "":
		results, err = s.CategoryRepository.FindAvailableByRestaurantGroupID(restaurantGroupID)
	case util.CATEGORY_TYPE_PRODUCT, util.CATEGORY_TYPE_INGREDIENT:
		results, err = s.CategoryRepository.FindAvailableByRestaurantGroupIDAndType(restaurantGroupID, cateType)
	default:
		return nil, exception.CreateError(exception.CodeValueInvalid, fmt.Sprintf("CateType: %s is invalid", cateType))
	}

	if err != nil {
		return nil, exception.CreateErrorWithRootCause(exception.CodeSystemError, fmt.Sprintf("Can not get Categories. restaurantGroupID=[%d], cateType=[%s]", restaurantGroupID, cateType), err)
	}

	return &results, nil
}
