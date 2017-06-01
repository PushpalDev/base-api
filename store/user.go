package store

import (
	"context"

	"github.com/pushpaldev/base-api/helpers/params"
	"github.com/pushpaldev/base-api/models"
)

func CreateUser(c context.Context, record *models.User) error {
	return FromContext(c).CreateUser(record)
}

func FindUserById(c context.Context, id string) (*models.User, error) {
	return FromContext(c).FindUserById(id)
}

func FindUser(c context.Context, params params.M) (*models.User, error) {
	return FromContext(c).FindUser(params)
}

func ActivateUser(c context.Context, activationKey string, id string) error {
	return FromContext(c).ActivateUser(activationKey, id)
}

func UpdateUser(c context.Context, params params.M) error {
	return FromContext(c).UpdateUser(Current(c), params)
}
