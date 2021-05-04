package persistence

import (
	"log"

	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"gopkg.in/gorp.v3"
)

type CommentRepositoryImpl struct {
	Table string
	DbMap *gorp.DbMap
}

func NewCommentRepository(dbMap *gorp.DbMap) repository.CommentRepository {
	return &CommentRepositoryImpl{Table: "comment", DbMap: dbMap}
}

func (r *CommentRepositoryImpl) FindAvailableByProductID(productID int64) (*entity.CommentList, error) {
	var items []entity.Comment
	_, err := r.DbMap.Select(&items, "SELECT * FROM "+r.Table+" WHERE product_id=? AND available=1", productID)

	if err == nil {
		return entity.NewCommentList(items), nil
	}

	log.Print(err)
	return nil, err

}
