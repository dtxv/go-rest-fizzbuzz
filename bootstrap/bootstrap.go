package bootstrap

import (
	"github.com/dtxv/go-rest-fizzbuzz/config"
	"github.com/dtxv/go-rest-fizzbuzz/controller"
	"github.com/dtxv/go-rest-fizzbuzz/docs"
	"github.com/dtxv/go-rest-fizzbuzz/domain"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/rs/zerolog/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
)

// Bootstrap struct
type Bootstrap struct {
	AppConfig *config.AppConfig
	Router    *gin.Engine
}

// Run app
func RunBootstrap() Bootstrap {
	app := Bootstrap{}
	app.AppConfig = config.GlobalConfig
	log.Logger = config.New(app.AppConfig.LoggerConfig)

	// services
	fizzBuzzService := domain.NewFizzBuzzService()

	// validators
	validator := domain.NewValidator()

	// handlers
	fizzBuzzHandler := controller.NewHandler(fizzBuzzService, validator)

	router := gin.New()
	gin.ForceConsoleColor()

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Use(logger.SetLogger(
		logger.WithSkipPath([]string{"/health"}),
		logger.WithUTC(true),
	))
	router.Use(cors.Default())
	router.Use(gin.Recovery())
	router.GET("/", controller.HealthCheck)
	router.GET("/health", controller.HealthCheck)
	appRouter := setRoutes(router, *fizzBuzzHandler)
	app.Router = appRouter
	return app
}
