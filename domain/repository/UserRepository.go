package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
)

type UserRepository interface {
	FindByUserName(userName string) (*entity.User, error)
	FindByUserNameAndPassword(userName string, password string) (*entity.User, error)
	FindAll() ([]entity.User, error)
}
