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

func (r *ProductRepositoryImpl) FindAvailableByID(id int64) (*entity.Product, error) {
	var item entity.Product
	return &item, r.DbMap.SelectOne(&item, "SELECT * FROM "+r.Table+" WHERE id=? AND available=1", id)
}
