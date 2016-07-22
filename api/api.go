package api

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/config"
	"github.com/iris-contrib/middleware/cors"
)

func NewServer() *iris.Framework {
	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Link", "Location", "Accept", "Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders: []string{"Link", "Location"},
	})

	server := iris.New(config.Iris{
		DisableBanner: true,
		Tester: config.Tester{ListeningAddr: "iris-go.com:1993/api/v1", ExplicitURL: false, Debug: false},
	})
	server.Use(crs)

	return server
}