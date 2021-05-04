package persistence

import (
	"github.com/smarest/smarest-api/domain/repository"
	"gopkg.in/gorp.v3"
)

type IngredientRepositoryImpl struct {
	Table string
	DbMap *gorp.DbMap
}

func NewIngredientRepository(dbMap *gorp.DbMap) repository.IngredientRepository {
	return &IngredientRepositoryImpl{Table: "ingredient", DbMap: dbMap}
}
