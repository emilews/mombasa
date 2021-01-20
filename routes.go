package main

import (
	"mombasa/handlers"
)

func initializeRoutes() {
	v1Routes := router.Group("/api/v1")
	{
		v1Routes.GET("/all", handlers.GetAllCryptos)
	}
	cryptoRoutes := router.Group("/api/c¡v1/crypto")
	{
		cryptoRoutes.GET("/", handlers.GetCrypto)
		cryptoRoutes.GET("/amount", handlers.GetCalculatedCryptoValue)
	}
	fiatRoutes := router.Group("/api/v1/fiat")
	{
		fiatRoutes.GET("/", handlers.GetCryptoWithSpecificFiat)
		fiatRoutes.GET("/amount")
	}
}
