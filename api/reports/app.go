package reports

import (
	"log"

	"github.com/kataras/iris"
)

func New(server *iris.Framework, params map[string]interface{}) {
	baseUrl := params["baseUrl"].(string)
	if val, ok := params["service"]; ok {
		service = val.(Reporter)
		handler := NewHandler(service)
		server.Get(baseUrl + "/reports", handler.Report)
	} else {
		log.Fatal("No service defined.")
	}
}