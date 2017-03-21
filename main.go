package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/dernise/base-api/server"
	"github.com/dernise/base-api/services"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
	"gopkg.in/mgo.v2"
	"os"
)

func main() {
	api := server.API{Router: gin.Default(), Config: viper.New()}

	err := api.LoadEnvVariables()
	if err != nil {
		panic(err)
	}

	err = api.SetupViper("prod")
	if err != nil {
		panic(err)
	}

	api.EmailSender = services.NewSendGridEmailSender(api.Config)

	session, err := mgo.Dial(os.Getenv("DB_HOST"))
	if err != nil {
		panic(err)
	}
	defer session.Close()
	api.Database = session.DB(os.Getenv("DB_NAME"))

	govalidator.SetFieldsRequiredByDefault(true)

	api.SetupRouter()
	api.Router.Run(api.Config.GetString("host.address"))
}
