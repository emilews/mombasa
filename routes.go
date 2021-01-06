package main

import (
	"mombasa/handlers"
)

func initializeRoutes() {
	apiRoutes := router.Group("/api/v1")
	{
		apiRoutes.GET("/all", handlers.GetAllCryptos)
		apiRoutes.GET("/fiat", handlers.GetCryptoWithSpecificFiat)
		apiRoutes.GET("/crypto", handlers.GetCrypto)
	}
}
