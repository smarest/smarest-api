package persistence

import (
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"gopkg.in/gorp.v3"
)

type AreaRepositoryImpl struct {
	Table string
	DbMap *gorp.DbMap
}

func NewAreaRepository(dbMap *gorp.DbMap) repository.AreaRepository {
	return &AreaRepositoryImpl{Table: "area", DbMap: dbMap}
}

func (r *AreaRepositoryImpl) FindByID(id int64) (*entity.Area, error) {
	var item entity.Area
	err := r.DbMap.SelectOne(&item, "SELECT * FROM "+r.Table+" WHERE id=?", id)

	if err == nil {
		return &item, nil
	}
	return nil, err
}
func (r *AreaRepositoryImpl) FindByRestaurantID(resId int64) ([]entity.Area, error) {
	var items []entity.Area
	_, err := r.DbMap.Select(&items, "SELECT * FROM "+r.Table+" WHERE restaurant_id=? AND available=1", resId)

	if err == nil {
		return items, nil
	}

	return nil, err
}

func (r *AreaRepositoryImpl) FindAll() ([]entity.Area, error) {
	var items []entity.Area
	_, err := r.DbMap.Select(&items, "SELECT * FROM "+r.Table)

	if err == nil {
		return items, nil
	}

	return nil, err

}
