package controller

import (
	"net/http"

	"github.com/dtxv/go-rest-fizzbuzz/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Handler struct
type Handler struct {
	fizzBuzzService domain.IFizzBuzzService
	validator       domain.IValidator
}

// NewClassifiedHandler return a new service for handlerclassified
func NewHandler(fizzBuzzService domain.IFizzBuzzService, validator domain.IValidator) *Handler {
	return &Handler{
		fizzBuzzService,
		validator,
	}
}

// Fizz Buzz handler
// @Summary      Fizz Buzz list
// @Description  Writes all numbers from 1 to limit by replacing all multiples of int1 by str1, all multiples of int2 by str2, and all multiples of int1 and int2 by str1+str2.
// @ID           fizz-buzz
// @Accept       text/plain
// @Produce      json
// @Param        int1   query      int     false  "The first number"
// @Param        int2   query      int     false  "The second number"
// @Param        limit  query      int     false  "The limit number"
// @Param        str1   query      string  false  "The first text"
// @Param        str2   query      string  false  "The second text"
// @Success      200    {integer}  int
// @Failure      400    {string}   string
// @Failure      500    {string}   string
// @Router       /fizz-buzz [get]
func (h Handler) FizzBuzz(c *gin.Context) {
	model, err := parseRequest(c)
	if err != nil {
		log.Warn().Err(err).Msg("Could not parse parameters")
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = h.validator.Validate(*model)
	if err != nil {
		log.Warn().Err(err).Msg("Could not validate parameters")
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	result := h.fizzBuzzService.FillFizzBuzz(*model)

	c.JSON(http.StatusOK, result)
}
