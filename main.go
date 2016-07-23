package main

import (
	"github.com/kataras/iris"
	"github.com/gustavohenrique/poupaniquel/api"
	"github.com/gustavohenrique/poupaniquel/api/database"
	"github.com/gustavohenrique/poupaniquel/api/webpage"
	"github.com/gustavohenrique/poupaniquel/api/transactions"
	"github.com/gustavohenrique/poupaniquel/api/reports"
	"github.com/gustavohenrique/poupaniquel/api/importers/nubank"
)

func main() {
	database.Create()

	baseUrl := "/api/v1"
	server := api.NewServer()
	
	webpage.New(server)
	transactions.New(server, params(baseUrl, transactions.NewService(transactions.NewDao())))
	reports.New(server, params(baseUrl, reports.NewService(reports.NewDao())))
	nubank.New(server, params(baseUrl, nubank.NewService(nubank.Origin)))

	banner := `
 ۜ\(סּںסּَ' )/ۜ
Poupaniquel API v0.0.1 Iris v` + iris.Version

	server.Logger.PrintBanner(banner, "\nVisit http://localhost:7000 in Google Chrome, Firefox, Opera or Safari.")
	server.Listen(":7000")
}

func params(baseUrl string, service interface{}) map[string]interface{} {
	return map[string]interface{}{
		"baseUrl": baseUrl,
		"service": service,
	}
}