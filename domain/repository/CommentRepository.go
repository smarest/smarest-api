package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type CommentRepository interface {
	FindByID(id int64) (*entity.Comment, error)
	FindAll() ([]entity.Comment, error)
}
