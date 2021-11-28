package controller

import (
	"fmt"

	"github.com/dtxv/go-rest-fizzbuzz/domain"

	"github.com/gin-gonic/gin"
)

// Parse Request
func parseRequest(c *gin.Context) (*domain.FizzBuzzModel, error) {
	var request FizzBuzzRequest

	if err := c.ShouldBindQuery(&request); err != nil {
		return nil, fmt.Errorf("impossible de binder les paramètres à l'objet SearchParams : %s", err)
	}

	model := &domain.FizzBuzzModel{
		Int1:  request.Int1,
		Int2:  request.Int2,
		Limit: request.Limit,
		Str1:  request.Str1,
		Str2:  request.Str2,
	}
	return model, nil
}
