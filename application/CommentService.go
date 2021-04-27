package application

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-api/domain/repository"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

type CommentService struct {
	CommentRepository repository.CommentRepository
	CommentFactory    entity.CommentFactory
}

func NewCommentService(CommentRepository repository.CommentRepository,
	CommentFactory entity.CommentFactory) *CommentService {
	return &CommentService{CommentRepository: CommentRepository,
		CommentFactory: CommentFactory}
}

func (s *CommentService) GetByID(c *gin.Context) {
	id, paramErr := strconv.ParseInt(c.Params.ByName("id"), 0, 64)
	if paramErr != nil {
		c.JSON(400, exception.CreateError(exception.CodeValueInvalid, "commentID invalid."))
		return
	}

	var ingredient, err = s.CommentRepository.FindByID(id)
	if err != nil {
		log.Print(err.Error())
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "comment not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, ingredient)
	} else {
		c.JSON(200, s.CommentFactory.Create(ingredient, fields))
	}
}

func (s *CommentService) GetAll(c *gin.Context) {
	var ingredients, err = s.CommentRepository.FindAll()
	if err != nil {
		c.JSON(404, exception.CreateError(exception.CodeNotFound, "comments not found."))
		return
	}

	var fields = c.Query("fields")
	if fields == "" {
		c.JSON(200, ingredients)
	} else {
		c.JSON(200, s.CommentFactory.CreateList(ingredients, fields))
	}

}
