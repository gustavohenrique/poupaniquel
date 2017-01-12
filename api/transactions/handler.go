package transactions

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

type Handler struct {}

var service TransactionManager
var helper = NewHelper()

func NewHandler(s TransactionManager) *Handler {
	service = s
	return &Handler{}
}

func (this *Handler) FetchAll(ctx *gin.Context) {
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
		log.Println("Error on FetchAll.", err)
	}
	ctx.Writer.Header().Set("link", link)
	ctx.JSON(status, transactions)
}

func (this *Handler) FetchOne(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err, transaction := service.FetchOne(id)
	status := 200
	if err != nil {
		log.Println("Error on FetchOne. ID =", id, err)
		status = 404
	}
	ctx.JSON(status, transaction)
}

func (this *Handler) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))
	err := service.Delete(id)
	status := 200
	if err != nil {
		log.Println("Error on Delete.", err)
		status = 500
	}
	ctx.JSON(status, err)
}

func (this *Handler) Create(ctx *gin.Context) {
	transaction, err := getPostBodyFrom(ctx)
	if err == nil {
		err = transaction.Validate()
	}
	status, response := save(transaction, err)
	ctx.JSON(status, response)
}

func (this *Handler) Update(ctx *gin.Context) {
	transaction, err := getPostBodyFrom(ctx)
	if err == nil {
		id, _ := strconv.Atoi(ctx.Params.ByName("id"))
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
		log.Println("Error on save.", err)
		status = 400
		response = helper.CreateErrorMap(err)
	}
	return status, response
}

func getPostBodyFrom(ctx *gin.Context) (Transaction, error) {
	transaction := Transaction{}
	// requestData := strings.NewReader(string(ctx.PostBody()))
	ctx.BindJSON(&transaction)
	// err := json.NewDecoder(requestData).Decode(&transaction)
	return transaction, nil
}
