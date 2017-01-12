package nubank

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

func New(server *gin.Engine, params map[string]interface{}) {
	baseUrl := params["baseUrl"].(string)
	if val, ok := params["service"]; ok {
		service = val.(ApiImporter)
		handler := NewHandler(service)
		server.GET(baseUrl+"/import/nubank", handler.Hello)
		server.POST(baseUrl+"/import/nubank", handler.ImportData)
	} else {
		log.Fatal("No service defined.")
	}
}
