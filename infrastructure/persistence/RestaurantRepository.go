package persistence

import (
	"strconv"
	"strings"

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

func (r *RestaurantRepositoryImpl) FindAll() ([]entity.Restaurant, error) {
	var items []entity.Restaurant
	_, err := r.DAOImpl.FindAll(&items)
	return items, err
}

func (r *RestaurantRepositoryImpl) FindProductByID(restaurantID int64, productID int64) (*entity.Product, error) {
	var item entity.Product
	return &item, r.DbMap.SelectOne(&item, "SELECT p.id,p.name,p.category_id,p.unit_id,p.description,p.image,rp.default_status,rp.price,rp.quantity_on_single_order,rp.available,rp.creator,rp.created_date,rp.updater,rp.last_updated_date FROM restaurant_product rp LEFT JOIN product p ON rp.product_id=p.id WHERE rp.restaurant_id=? AND p.id=?", restaurantID, productID)
}

func (r *RestaurantRepositoryImpl) FindProductByIDs(restaurantID int64, productIDs []int64) (entity.ProductList, error) {
	if len(productIDs) == 0 {
		return entity.CreateEmptyProductList(), nil
	}
	var items []entity.Product
	productIDString := make([]string, len(productIDs))
	for i, productID := range productIDs {
		productIDString[i] = strconv.FormatInt(productID, 10)
	}
	_, err := r.DAOImpl.Select(&items, "SELECT p.id,p.name,p.category_id,p.unit_id,p.description,p.image,rp.default_status,rp.price,rp.quantity_on_single_order,rp.available,rp.creator,rp.created_date,rp.updater,rp.last_updated_date FROM restaurant_product rp LEFT JOIN product p ON rp.product_id=p.id WHERE rp.restaurant_id=? AND p.id IN ("+strings.Join(productIDString, ",")+")", restaurantID)
	return entity.NewProductList(items), err
}
