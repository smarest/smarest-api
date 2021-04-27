package persistence

import (
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/infrastructure/persistence"
	"gopkg.in/gorp.v3"
)

type ProductRepositoryImpl struct {
	*persistence.DAOImpl
}

func NewProductRepository(dbMap *gorp.DbMap) repository.ProductRepository {
	return &ProductRepositoryImpl{persistence.NewDAOImpl("`product`", dbMap)}
}

func (r *ProductRepositoryImpl) FindByID(id int64) (*entity.Product, error) {
	var item entity.Product
	return &item, r.DAOImpl.FindByID(id, &item)
}
func (r *ProductRepositoryImpl) FindByRestaurantID(resId int64) ([]entity.Product, error) {
	var products []entity.Product
	_, err := r.DbMap.Select(&products, "SELECT p.id,p.name,p.category_id,p.unit_id,p.description,p.image,rp.default_status,rp.price,rp.quantity_on_single_order,rp.available,rp.creator,rp.created_date,rp.updater,rp.last_updated_date FROM restaurant_product rp LEFT JOIN product p ON p.id=rp.product_id WHERE rp.restaurant_id=? AND rp.available=1 AND p.available=1", resId)

	if err == nil {
		return products, nil
	}

	return nil, err
}

func (r *ProductRepositoryImpl) FindByCategoryID(cateID int64) ([]entity.Product, error) {
	var products []entity.Product
	_, err := r.DAOImpl.Select(&products, "SELECT p.id,p.name,p.category_id,p.unit_id,p.description,p.image,rp.default_status,rp.price,rp.quantity_on_single_order,rp.available,rp.creator,rp.created_date,rp.updater,rp.last_updated_date FROM restaurant_product rp LEFT JOIN product p ON p.id=rp.product_id WHERE p.category_id=?", cateID)

	if err == nil {
		return products, nil
	}

	return nil, err
}

func (r *ProductRepositoryImpl) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	_, err := r.DAOImpl.Select(&products, "SELECT * FROM "+r.Table)

	if err == nil {
		return products, nil
	}

	return nil, err

}
