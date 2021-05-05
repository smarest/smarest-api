package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type TableRepository interface {
	FindAvailableByAreaID(areaID int64) (*entity.TableList, error)
	FindAvailableByRestaurantIDAndIDs(restaurantID int64, tableIds []int64) (*entity.TableList, error)
	FindByIDs(tableIds []int64) (*entity.TableList, error)
}
