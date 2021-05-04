package persistence

import (
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/infrastructure/persistence"
	"gopkg.in/gorp.v3"
)

type AreaRepositoryImpl struct {
	*persistence.DAOImpl
}

func NewAreaRepository(dbMap *gorp.DbMap) repository.AreaRepository {
	return &AreaRepositoryImpl{persistence.NewDAOImpl("`area`", dbMap)}
}

func (r *AreaRepositoryImpl) FindByID(id int64) (*entity.Area, error) {
	var item entity.Area
	return &item, r.DAOImpl.FindByID(id, &item)
}

func (r *AreaRepositoryImpl) FindAvailableByID(id int64) (*entity.Area, error) {
	var item entity.Area
	return &item, r.DbMap.SelectOne(&item, "SELECT * FROM "+r.Table+" WHERE id=? AND available=1", id)
}

func (r *AreaRepositoryImpl) FindAvailableByRestaurantID(resId int64) (*entity.AreaList, error) {
	var items []entity.Area
	_, err := r.DAOImpl.Select(&items, "SELECT * FROM "+r.Table+" WHERE restaurant_id=? AND available=1", resId)
	return entity.NewAreaList(items), err
}
