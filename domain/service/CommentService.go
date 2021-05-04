package service

import (
	"database/sql"
	"fmt"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type CommentService struct {
	ProductRepository repository.ProductRepository
	CommentRepository repository.CommentRepository
}

func NewCommentService(productRepository repository.ProductRepository, commentRepository repository.CommentRepository) *CommentService {
	return &CommentService{productRepository, commentRepository}
}

func (s *CommentService) FindAvailableByRestaurantGroupIDAndProductID(restaurantGroupID int64, productID int64) (*entity.CommentList, *exception.Error) {
	product, err := s.ProductRepository.FindAvailableByID(productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.CreateErrorWithRootCause(exception.CodeNotFound, fmt.Sprintf("ProductID=[%d] not found.", productID), err)
		}
		return nil, exception.CreateErrorWithRootCause(exception.CodeSystemError, fmt.Sprintf("Can not get ProductID=[%d].", productID), err)

	}

	if product.RestaurantGroupID != restaurantGroupID {
		return nil, exception.CreateError(exception.CodeValueInvalid, fmt.Sprintf("ProductID=[%d] not belong restaurantGroupID=[%d]", productID, restaurantGroupID))
	}

	results, err := s.CommentRepository.FindAvailableByProductID(productID)
	if err != nil {
		return nil, exception.CreateErrorWithRootCause(exception.CodeSystemError, "Can not get Comments.", err)
	}

	return results, nil
}
