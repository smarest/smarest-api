package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type ProductRepository interface {
	FindByID(id int64) (*entity.Product, error)
	FindAvailableByID(id int64) (*entity.Product, error)
}
