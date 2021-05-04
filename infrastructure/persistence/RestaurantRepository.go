package persistence

import (
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/infrastructure/persistence"
	"gopkg.in/gorp.v3"
)

type RestaurantRepositoryImpl struct {
	*persistence.DAOImpl
}

func NewRestaurantRepository(dbMap *gorp.DbMap) repository.RestaurantRepository {
	return &RestaurantRepositoryImpl{persistence.NewDAOImpl("`restaurant`", dbMap)}
}

func (r *RestaurantRepositoryImpl) FindByID(id int64) (*entity.Restaurant, error) {
	var item entity.Restaurant
	return &item, r.DAOImpl.FindByID(id, &item)
}

func (r *RestaurantRepositoryImpl) FindAvailableByID(id int64) (*entity.Restaurant, error) {
	var item entity.Restaurant
	return &item, r.DbMap.SelectOne(&item, "SELECT * FROM "+r.Table+" WHERE id=? AND available=1", id)
}

func (r *RestaurantRepositoryImpl) FindByRestaurantGroupID(restaurantGroupID int64) (*entity.RestaurantList, error) {
	var items []entity.Restaurant
	_, err := r.DAOImpl.Select(&items, "SELECT * FROM "+r.Table+" WHERE restaurant_group_id=?", restaurantGroupID)
	return entity.NewRestaurantList(items), err
}

func (r *RestaurantRepositoryImpl) FindAvailableByRestaurantGroupID(restaurantGroupID int64) (*entity.RestaurantList, error) {
	var items []entity.Restaurant
	_, err := r.DAOImpl.Select(&items, "SELECT * FROM "+r.Table+" WHERE restaurant_group_id=? AND available=1", restaurantGroupID)
	return entity.NewRestaurantList(items), err
}

func (r *RestaurantRepositoryImpl) FindByCode(code string) (*entity.Restaurant, error) {
	var item entity.Restaurant
	return &item, r.DAOImpl.SelectOne(&item, "SELECT * FROM "+r.Table+" WHERE code=?", code)
}

func (r *RestaurantRepositoryImpl) FindAll() (*entity.RestaurantList, error) {
	var items []entity.Restaurant
	_, err := r.DAOImpl.FindAll(&items)
	return entity.NewRestaurantList(items), err
}
