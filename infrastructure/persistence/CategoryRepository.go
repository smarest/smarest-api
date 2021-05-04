package persistence

import (
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/infrastructure/persistence"
	"gopkg.in/gorp.v3"
)

type CategoryRepositoryImpl struct {
	*persistence.DAOImpl
}

func NewCategoryRepository(dbMap *gorp.DbMap) repository.CategoryRepository {
	return &CategoryRepositoryImpl{persistence.NewDAOImpl("`category`", dbMap)}
}

func (r *CategoryRepositoryImpl) FindAvailableByRestaurantGroupID(groupID int64) (entity.CategoryList, error) {
	var items []entity.Category
	_, err := r.DAOImpl.Select(&items, "SELECT * FROM "+r.Table+" WHERE restaurant_group_id=? AND available=1", groupID)

	return entity.NewCategoryList(items), err
}

func (r *CategoryRepositoryImpl) FindAvailableByRestaurantGroupIDAndType(groupID int64, cateType string) (entity.CategoryList, error) {
	var items []entity.Category
	_, err := r.DAOImpl.Select(&items, "SELECT * FROM "+r.Table+" WHERE restaurant_group_id=? AND type=? AND available=1", groupID, cateType)

	return entity.NewCategoryList(items), err
}
