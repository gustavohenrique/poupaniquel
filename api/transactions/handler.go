package transactions

import (
	"fmt"
	"log"
	"encoding/json"
	"strings"
	"github.com/kataras/iris"
)

type Handler struct {}

var service TransactionManager
var helper = NewHelper()

func NewHandler(s TransactionManager) *Handler {
	service = s
	return &Handler{}
}

func (this *Handler) FetchAll(ctx *iris.Context) {
	params := helper.GetPageParameters(ctx)
	page := params["page"].(int)
	err, transactions := service.FetchAll(params)
	if len(transactions) == 0 {
		log.Println("No transactions found in page", page)
		transactions = []Transaction{}
	}
	status := 500
	var link string
	if err == nil {
		status = 200
		previous := page - 1
		if previous < 1 {
			previous = 1
		}
		next := page + 1
		link = fmt.Sprintf(`</api/v1/transactions?page=%d>; rel="previous", </api/v1/transactions?page=%d>; rel="next"`, previous, next)
	} else {
		log.Println("Error in FetchAll.", err)
	}
	ctx.SetHeader("link", link)
	ctx.JSON(status, transactions)
}

func (this *Handler) FetchOne(ctx *iris.Context) {
	id, _ := ctx.ParamInt("id")
	err, transaction := service.FetchOne(id)
	status := 200
	if err != nil {
		log.Println("Error in FetchOne. ID =", id, err)
		status = 404
	}
	ctx.JSON(status, transaction)
}

func (this *Handler) Delete(ctx *iris.Context) {
	id, _ := ctx.ParamInt("id")
	err := service.Delete(id)
	status := 204
	if err != nil {
		log.Println("Error in Delete.", err)
		status = 500
	}
	ctx.JSON(status, nil)
}

func (this *Handler) Create(ctx *iris.Context) {
	transaction, err := getPostBodyFrom(ctx)
	if err == nil {
		err = transaction.Validate()
	}
	status, response := save(transaction, err)
	ctx.JSON(status, response)
}

func (this *Handler) Update(ctx *iris.Context) {
	transaction, err := getPostBodyFrom(ctx)
	if err == nil {
		id, _ := ctx.ParamInt("id")
		transaction.Id = int64(id)
		err = transaction.Validate()
	}
	status, response := save(transaction, err)
	ctx.JSON(status, response)
}

func save(transaction Transaction, err error) (status int, response interface{}) {
	if err == nil {
		id := int64(0)
		err, id = service.Save(transaction)
		if err == nil {
			status = 200
			response = map[string]int64{"id": id}
		}
	}
	if err != nil {
		log.Println("Error in save.", err)
		status = 400
		response = helper.CreateErrorMap(err)
	}
	return status, response
}

func getPostBodyFrom(ctx *iris.Context) (Transaction, error) {
	transaction := Transaction{}
	requestData := strings.NewReader(string(ctx.PostBody()))
	err := json.NewDecoder(requestData).Decode(&transaction)
	return transaction, err
}
