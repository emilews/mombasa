package handlers

import (
	"mombasa/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCryptoWithSpecificFiat(c *gin.Context) {
	cryptoTicker := c.PostForm("crypto_ticker")
	fiatTicker := c.PostForm("fiat_ticker")
	c.JSON(http.StatusOK, models.GetCryptoWithSpecificFiat(fiatTicker, cryptoTicker))
}

func GetCrypto(c *gin.Context) {
	cryptoTicker := c.PostForm("crypto_ticker")
	c.JSON(http.StatusOK, models.GetCrypto(cryptoTicker))
}

func GetAllCryptos(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllCryptos())
}

func GetCalculatedCryptoValue(c *gin.Context){
	cryptoTicker := c.PostForm("crypto_ticker")
	cryptoAmount := c.PostForm("crypto_amount")
	c.JSON(http.StatusOK, models.GetCalculatedCryptoValue(cryptoTicker, cryptoAmount))
}