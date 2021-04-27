package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type IngredientRepository interface {
	FindByID(id int64) (*entity.Ingredient, error)
	FindAll() ([]entity.Ingredient, error)
}
