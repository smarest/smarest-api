package persistence

import (
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"gopkg.in/gorp.v3"
)

type CategoryRepositoryImpl struct {
	Table string
	DbMap *gorp.DbMap
}

func NewCategoryRepository(dbMap *gorp.DbMap) repository.CategoryRepository {
	return &CategoryRepositoryImpl{Table: "category", DbMap: dbMap}
}

func (r *CategoryRepositoryImpl) FindByID(id int64) (*entity.Category, error) {
	var item entity.Category
	err := r.DbMap.SelectOne(&item, "SELECT * FROM "+r.Table+" WHERE id=?", id)

	if err == nil {
		return &item, nil
	}
	return nil, err
}
func (r *CategoryRepositoryImpl) FindByType(cateType string) (entity.CategoryList, error) {
	var items []entity.Category
	_, err := r.DbMap.Select(&items, "SELECT * FROM "+r.Table+" WHERE type=?", cateType)

	return entity.NewCategoryList(items), err
}

func (r *CategoryRepositoryImpl) FindAll() (entity.CategoryList, error) {
	var items []entity.Category
	_, err := r.DbMap.Select(&items, "SELECT * FROM "+r.Table)

	return entity.NewCategoryList(items), err

}
