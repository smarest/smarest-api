package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type CategoryRepository interface {
	FindAvailableByRestaurantGroupID(groupID int64) (entity.CategoryList, error)
	FindAvailableByRestaurantGroupIDAndType(groupID int64, cateType string) (entity.CategoryList, error)
}
