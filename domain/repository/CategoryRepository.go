package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type CategoryRepository interface {
	FindByID(id int64) (*entity.Category, error)
	FindByType(cateType string) (entity.CategoryList, error)
	FindAll() (entity.CategoryList, error)
}
