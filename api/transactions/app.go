package transactions

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

func New(server *gin.Engine, params map[string]interface{}) {
	baseUrl := params["baseUrl"].(string)
	if val, ok := params["service"]; ok {
		service = val.(TransactionManager)
		handler := NewHandler(service)
		server.GET(baseUrl + "/transactions", handler.FetchAll)
		server.GET(baseUrl + "/transactions/:id", handler.FetchOne)
		server.POST(baseUrl + "/transactions", handler.Create)
		server.PUT(baseUrl + "/transactions/:id", handler.Update)
		server.DELETE(baseUrl + "/transactions/:id", handler.Delete)
	} else {
		log.Fatal("No service defined.")
	}
}