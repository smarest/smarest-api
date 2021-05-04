package persistence

import (
	"strconv"
	"strings"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/infrastructure/persistence"
	"gopkg.in/gorp.v3"
)

type RestaurantProductRepositoryImpl struct {
	*persistence.DAOImpl
}

func NewRestaurantProductRepository(dbMap *gorp.DbMap) repository.RestaurantProductRepository {
	return &RestaurantProductRepositoryImpl{persistence.NewDAOImpl("`restaurant_product`", dbMap)}
}

func (r *RestaurantProductRepositoryImpl) FindAvailableProductsByCategoryID(restaurantID int64, cateID int64) (*entity.ProductList, error) {
	var products []entity.Product
	_, err := r.DAOImpl.Select(&products, "SELECT p.id,p.restaurant_group_id,p.name,p.category_id,p.unit_id,p.description,p.image,rp.default_status,rp.price,rp.quantity_on_single_order,rp.available,rp.creator,rp.created_date,rp.updater,rp.last_updated_date FROM restaurant_product rp LEFT JOIN product p ON p.id=rp.product_id WHERE rp.restaurant_id=? AND p.category_id=? AND rp.available=1 AND p.available=1", restaurantID, cateID)

	if err == nil {
		return entity.NewProductList(products), nil
	}

	return nil, err
}

func (r *RestaurantProductRepositoryImpl) FindAvailableProductByID(restaurantID int64, productID int64) (*entity.Product, error) {
	var item entity.Product
	return &item, r.DbMap.SelectOne(&item, "SELECT p.id,p.name,p.restaurant_group_id,p.category_id,p.unit_id,p.description,p.image,rp.default_status,rp.price,rp.quantity_on_single_order,rp.available,rp.creator,rp.created_date,rp.updater,rp.last_updated_date FROM restaurant_product rp LEFT JOIN product p ON rp.product_id=p.id WHERE rp.restaurant_id=? AND p.id=? AND p.available=1 AND rp.available=1", restaurantID, productID)
}

func (r *RestaurantProductRepositoryImpl) FindAvailableProductsByRestaurantIDAndIDs(restaurantID int64, productIDs []int64) (*entity.ProductList, error) {
	if len(productIDs) == 0 {
		return entity.CreateEmptyProductList(), nil
	}
	var items []entity.Product
	productIDString := make([]string, len(productIDs))
	for i, productID := range productIDs {
		productIDString[i] = strconv.FormatInt(productID, 10)
	}
	_, err := r.DAOImpl.Select(&items, "SELECT p.id,p.name,p.restaurant_group_id,p.category_id,p.unit_id,p.description,p.image,rp.default_status,rp.price,rp.quantity_on_single_order,rp.available,rp.creator,rp.created_date,rp.updater,rp.last_updated_date FROM restaurant_product rp LEFT JOIN product p ON rp.product_id=p.id WHERE rp.restaurant_id=? AND p.id IN ("+strings.Join(productIDString, ",")+") AND p.available=1 AND rp.available=1", restaurantID)
	return entity.NewProductList(items), err
}
