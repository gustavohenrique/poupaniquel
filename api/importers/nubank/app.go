package nubank

import (
	"log"

	"github.com/kataras/iris"
)

func New(server *iris.Framework, params map[string]interface{}) {
	baseUrl := params["baseUrl"].(string) + "/import"
	if val, ok := params["service"]; ok {
		service = val.(ApiImporter)
		handler := NewHandler(service)
		server.Post(baseUrl + "/nubank", handler.ImportData)
	} else {
		log.Fatal("No service defined.")
	}
}