package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type TableRepository interface {
	FindByID(id int64) (*entity.Table, error)
	FindByAreaID(areaID int64) ([]entity.Table, error)
	FindByIDs(areaIDs []int64) (entity.TableList, error)
	FindAll() ([]entity.Table, error)
}
