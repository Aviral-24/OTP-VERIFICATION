package api

import (
	"context"

	"go-sms-verify-yt/data"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func (app *Config) Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		app.writeJSON(c, http.StatusOK, "Server is running")
	}
}

func (app *Config) SendSMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.OTPData
		defer cancel()

		err := app.validateBody(c, &payload)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		newData := data.OTPData{
			PhoneNumber: payload.PhoneNumber,
		}
		_, err = app.twilioSendOTP(newData.PhoneNumber)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "OTP sent successfully")
	}
}

func (app *Config) VerifySMS() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), appTimeout)
		var payload data.VerifyData
		defer cancel()

		err := app.validateBody(c, &payload)
		if err != nil {
			app.errorJSON(c, err)
			return
		}

		newData := data.VerifyData{
			PhoneNumber: payload.PhoneNumber,
			Code:        payload.Code,
		}
		_, err = app.twilioVerifyOTP(newData.PhoneNumber, newData.Code)
		if err != nil {
			app.errorJSON(c, err)
			return
		}
		app.writeJSON(c, http.StatusAccepted, "OTP verified successfully")
	}
}
