package service

import (
	"github.com/smarest/smarest-api/domain/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{productRepository}
}
