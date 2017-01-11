package api

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
)

func NewServer() *iris.Framework {
	/*crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Link", "Location", "Accept", "Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders: []string{"Link", "Location"},
	})*/
	iris.Use(cors.Default())
	server := iris.New(iris.Configuration{
		DisableBanner: true,
		// Tester: iris.Default.Config.Tester{ListeningAddr: "iris-go.com:1993/api/v1", ExplicitURL: false, Debug: false},
	})
	//iris.Use(crs)

	return server
}