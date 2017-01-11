package main

import (
	"net"
	"log"
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"

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
Poupaniquel API v0.0.3 Gin ` + gin.Version

	ip := getIpAddress()
	msg := fmt.Sprintf("%s\nVisit http://%s:7000 in Google Chrome, Firefox, Opera or Safari.", banner, ip)
	log.Println(msg)
	server.Run(":7000")
}

func params(baseUrl string, service interface{}) map[string]interface{} {
	return map[string]interface{}{
		"baseUrl": baseUrl,
		"service": service,
	}
}

func getIpAddress() string {
	ifaces, _ := net.Interfaces()

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String()
		}
	}
	return "localhost"
}