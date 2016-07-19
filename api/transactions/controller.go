package transactions

import (
	"github.com/kataras/iris"
	"fmt"
	"log"
	"encoding/json"
	"strings"
)

type Controller struct {}

var service = NewService()

func NewController(base string) {
	api := &Controller{}
	iris.Get(base + "/transactions", api.FetchAll)
	iris.Get(base + "/transactions/:id", api.FetchOne)
	iris.Post(base + "/transactions", api.Create)
	iris.Put(base + "/transactions/:id", api.Update)
	iris.Delete(base + "/transactions/:id", api.Delete)
}

func (this *Controller) FetchAll(ctx *iris.Context) {
	params := this.getParamsFrom(ctx)
	page := params["page"].(int)
	err, transactions := service.FetchAll(params)
	if len(transactions) == 0 {
		log.Println("No transactions found in page", page)
		transactions = []Transaction{}
	}
	status := 500
	link := ""
	if err == nil {
		status = 200
		previous := page - 1
		if previous < 1 {
			previous = 1
		}
		next := page + 1
		link = fmt.Sprintf(`</api/v1/transactions?page=%d>; rel="previous", </api/v1/transactions?page=%d>; rel="next"`, previous, next)
	}
	ctx.SetHeader("link", link)
	ctx.JSON(status, transactions)
}

func (this *Controller) FetchOne(ctx *iris.Context) {
	id, _ := ctx.ParamInt("id")
	err, transaction := service.FetchOne(id)
	status := 200
	if err != nil {
		log.Println("Error FetchOne id=", id, err)
		status = 404
	}
	ctx.JSON(status, transaction)
}

func (this *Controller) Delete(ctx *iris.Context) {
	id, _ := ctx.ParamInt("id")
	err := service.Delete(id)
	status := 204
	if err != nil {
		log.Println("Error Delete.", err)
		status = 500
	}
	ctx.JSON(status, nil)
}

func (this *Controller) Create(ctx *iris.Context) {
	transaction := Transaction{}
	requestData := strings.NewReader(string(ctx.PostBody()))
	err := json.NewDecoder(requestData).Decode(&transaction)
	if err == nil {
		err = transaction.isValid()
	}
	status, response := this.save(err, transaction)
	ctx.JSON(status, response)
}

func (this *Controller) Update(ctx *iris.Context) {
	transaction := Transaction{}
	requestData := strings.NewReader(string(ctx.PostBody()))
	err := json.NewDecoder(requestData).Decode(&transaction)
	if err == nil {
		id, _ := ctx.ParamInt("id")
		transaction.Id = int64(id)
		err = transaction.isValid()
	}
	status, response := this.save(err, transaction)
	ctx.JSON(status, response)
}

func (this *Controller) save(err error, transaction Transaction) (status int, response interface{}) {
	if err == nil {
		id := int64(0)
		err, id = service.Save(transaction)
		if err == nil {
			status = 200
			response = map[string]int64{"id": id}
		}
	}
	if err != nil {
		log.Println("Error saving.", err)
		status = 400
		response = this.getError(err)
	}
	return status, response
}

func (*Controller) getError(err error) map[string]string {
	log.Println(err)
	return map[string]string{
		"code": "BadRequest",
		"message": fmt.Sprintf("%v", err),
	}
}

func (*Controller) getParamsFrom(ctx *iris.Context) map[string]interface{} {
	var page, perPage = 1, 1
	page, err := ctx.URLParamInt("page")
	if err != nil || page < 1 {
		page = 1
	}
	perPage, err = ctx.URLParamInt("perPage")
	if err != nil || perPage < 1 {
		perPage = 1
	}
	sort := ctx.URLParam("sort")
	if sort == "" {
		sort = "id"
	}
	return map[string]interface{}{
		"page": page,
		"perPage": perPage,
		"sort": sort,
	}
}