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

func (r *CommentRepositoryImpl) FindByID(id int64) (*entity.Comment, error) {
	var Comment entity.Comment
	err := r.DbMap.SelectOne(&Comment, "SELECT * FROM "+r.Table+" WHERE id=?", id)

	if err == nil {
		return &Comment, nil
	}
	return nil, err
}
func (r *CommentRepositoryImpl) FindAll() ([]entity.Comment, error) {
	var Comments []entity.Comment
	_, err := r.DbMap.Select(&Comments, "SELECT * FROM "+r.Table)

	if err == nil {
		return Comments, nil
	}

	log.Print(err)
	return nil, err

}
