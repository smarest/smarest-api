package service

import (
	"github.com/smarest/smarest-api/domain/repository"
)

type IngredientService struct {
	IngredientRepository repository.IngredientRepository
}

func NewIngredientService(ingredientRepository repository.IngredientRepository) *IngredientService {
	return &IngredientService{ingredientRepository}
}
