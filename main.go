package main

import (
	"github.com/kataras/iris"
	"github.com/gustavohenrique/poupaniquel/api"
	"github.com/gustavohenrique/poupaniquel/api/database"
	"github.com/gustavohenrique/poupaniquel/api/webpage"
	"github.com/gustavohenrique/poupaniquel/api/transactions"
	"github.com/gustavohenrique/poupaniquel/api/reports"
)

func main() {
	database.Create()

	baseUrl := "/api/v1"
	server := api.NewServer()
	
	webpage.New(server)
	transactions.New(server, map[string]interface{}{
		"baseUrl": baseUrl,
		"service": transactions.NewService(transactions.NewDao()),
	})
	reports.New(server, map[string]interface{}{
		"baseUrl": baseUrl,
		"service": reports.NewService(reports.NewDao()),
	})

	banner := `
 ۜ\(סּںסּَ' )/ۜ
Poupaniquel API v0.0.1 Iris v` + iris.Version

	server.Logger.PrintBanner(banner, "\nVisit http://localhost:7000 in Google Chrome, Firefox, Opera or Safari.")
	server.Listen(":7000")
}


