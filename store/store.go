package store

import (
	"github.com/pushpaldev/base-api/helpers/params"
	"github.com/pushpaldev/base-api/models"
)

type Store interface {
	CreateUser(*models.User) error
	FindUserById(string) (*models.User, error)
	ActivateUser(string, string) error
	FindUser(params.M) (*models.User, error)
	UpdateUser(*models.User, params.M) error
	AddLoginToken(*models.User, string) (*models.LoginToken, error)
	RemoveLoginToken(*models.User, string) error
}
