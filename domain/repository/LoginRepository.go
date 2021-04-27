package repository

import (
	"github.com/smarest/smarest-api/domain/entity"
	"github.com/smarest/smarest-common/domain/entity/exception"
)

// LoginRepository is used for get user information
type LoginRepository interface {
	GetUserByCookie(cookie string) (*entity.User, *exception.Error)
	GetUserByUserNameAndPassword(userName string, password string) (*entity.User, *exception.Error)
}
