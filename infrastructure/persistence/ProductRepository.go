package persistence

import (
	"fmt"
	"strings"

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

func (r *ProductRepositoryImpl) FindAvailableByID(id int64) (*entity.Product, error) {
	var item entity.Product
	return &item, r.DbMap.SelectOne(&item, "SELECT * FROM "+r.Table+" WHERE id=? AND available=1", id)
}
func (r *ProductRepositoryImpl) FindByIDs(productIDs []int64) (*entity.ProductList, error) {
	if len(productIDs) == 0 {
		return entity.CreateEmptyProductList(), nil
	}
	var items []entity.Product
	productIDStrs := make([]string, len(productIDs))
	for i, productID := range productIDs {
		productIDStrs[i] = fmt.Sprint(productID)
	}
	_, err := r.DbMap.Select(&items, "SELECT * FROM "+r.Table+" WHERE id IN("+strings.Join(productIDStrs, ",")+")")
	return entity.NewProductList(items), err

}
