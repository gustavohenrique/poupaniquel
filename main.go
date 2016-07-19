package main

import (
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
	"github.com/gustavohenrique/poupaniquel/api/database"
	"github.com/gustavohenrique/poupaniquel/api/transactions"
	"github.com/gustavohenrique/poupaniquel/api/reports"
	"github.com/gustavohenrique/poupaniquel/api/webpage"
)

func main() {
	database.Create()

	crs := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Link", "Location", "Accept", "Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders: []string{"Link", "Location"},
	})
	iris.Use(crs)

	base := "/api/v1"
	webpage.NewController()
	transactions.NewController(base)
	reports.NewController(base)

	iris.Config.DisableBanner = true
	banner := `
 ۜ\(סּںסּَ' )/ۜ
Poupaniquel API v0.0.1 Iris v` + iris.Version

	iris.Logger.PrintBanner(banner, "\nVisit http://localhost:7000 in Google Chrome, Firefox, Opera or Safari.")
	iris.Listen(":7000")
}
