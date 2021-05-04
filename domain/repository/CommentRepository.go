package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type CommentRepository interface {
	FindAvailableByProductID(productID int64) (*entity.CommentList, error)
}
