package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mombasa/models"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)
	models.MongoStore.Session = models.InitializeMongo()
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "content-type"},
	}))
	initializeRoutes()
	router.Run(":80")
}
