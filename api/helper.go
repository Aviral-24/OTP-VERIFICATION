package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(c *gin.Context, data any) error {
	err := c.BindJSON(data)
	if err != nil {
		return err
	}
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}

func (app *Config) writeJSON(c *gin.Context, status int, message string, data ...any) {
	respData := data
	if len(data) == 0 {
		respData = nil
	}
	c.JSON(status, jsonResponse{
		Status:  true,
		Message: message,
		Data:    respData,
	})
}

func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	c.JSON(statusCode, jsonResponse{
		Status:  false,
		Message: err.Error(),
	})
}
