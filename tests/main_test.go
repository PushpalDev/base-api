package tests

import (
	"os"
	"testing"

	"github.com/pushpaldev/base-api/models"
	"github.com/pushpaldev/base-api/server"
)

var api *server.API
var user *models.User
var authToken string

func TestMain(m *testing.M) {
	api = SetupApi()
	user, authToken = CreateUserAndGenerateToken()
	retCode := m.Run()
	api.Database.Session.Close()
	os.Exit(retCode)
}
