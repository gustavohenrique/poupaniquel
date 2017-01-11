package api

import (
	"time"

	"gopkg.in/gin-contrib/cors.v1"
	"gopkg.in/gin-gonic/gin.v1"
)

func NewServer() *gin.Engine {
	/*crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Link", "Location", "Accept", "Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders: []string{"Link", "Location"},
	})*/
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.Use(cors.New(cors.Config{
        // AllowOrigins:     []string{"*"},
        AllowAllOrigins:  true,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Link", "Location", "Accept", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Link", "Location"},
        MaxAge: 1 * time.Hour,
    }))

	return server
}