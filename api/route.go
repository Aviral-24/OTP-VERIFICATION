package api

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	Router *gin.Engine
}

func (app *Config) Routes() {
	app.Router.GET("/", app.Health())
	app.Router.POST("/send-otp", app.SendSMS())
	app.Router.POST("/verify-otp", app.VerifySMS())
}
