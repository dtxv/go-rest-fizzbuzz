package bootstrap

import (
	"github.com/dtxv/go-rest-fizzbuzz/controller"
	"github.com/gin-gonic/gin"
)

func setRoutes(router *gin.Engine, handler controller.Handler) *gin.Engine {

	router.GET("/fizz-buzz", handler.FizzBuzz)

	return router
}
