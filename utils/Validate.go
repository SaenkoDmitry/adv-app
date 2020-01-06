package utils

import (
	"github.com/SaenkoDmitry/advertisement-app/validations"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"net/http"
)

func ValidateAndBind(c *gin.Context, params interface{}) bool {
	if err := c.ShouldBind(params); err != nil {
		var errors []string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errors = append(errors, validations.FieldError{Err: fieldErr}.String())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errors,
		})
		return false
	}
	return true
}
