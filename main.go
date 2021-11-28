package main

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/dtxv/go-rest-fizzbuzz/bootstrap"
)

// @title           fizzbuzz
// @version         1.0
// @description     This API provides number transformations regarding specific rules
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://contact
// @contact.email  a@email.com

// @license.name  licence name
// @license.url   http://licence.url

// @host      localhost:5000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {
	app := bootstrap.RunBootstrap()
	err := app.Router.Run(fmt.Sprintf("%s:%d", app.AppConfig.HostAddress, app.AppConfig.HTTPPort))
	if err != nil {
		log.Fatal().Err(err).Msg("Application launch failure")
	}
}
