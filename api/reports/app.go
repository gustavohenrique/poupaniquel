package reports

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

func New(server *gin.Engine, params map[string]interface{}) {
	baseUrl := params["baseUrl"].(string)
	if val, ok := params["service"]; ok {
		service = val.(Reporter)
		handler := NewHandler(service)
		server.GET(baseUrl+"/reports", handler.Report)
	} else {
		log.Fatal("No service defined.")
	}
}
