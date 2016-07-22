package transactions

import (
	"log"

	"github.com/kataras/iris"
)

func New(server *iris.Framework, params map[string]interface{}) {
	baseUrl := params["baseUrl"].(string)
	if val, ok := params["service"]; ok {
		service = val.(TransactionManager)
		handler := NewHandler(service)
		server.Get(baseUrl + "/transactions", handler.FetchAll)
		server.Get(baseUrl + "/transactions/:id", handler.FetchOne)
		server.Post(baseUrl + "/transactions", handler.Create)
		server.Put(baseUrl + "/transactions/:id", handler.Update)
		server.Delete(baseUrl + "/transactions/:id", handler.Delete)
	} else {
		log.Fatal("No service defined.")
	}
}