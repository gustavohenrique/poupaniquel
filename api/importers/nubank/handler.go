package nubank

import (
	"log"
	"strings"
	"encoding/json"
	"time"

	"github.com/kataras/iris"

	"github.com/gustavohenrique/poupaniquel/api/transactions"
)

type Handler struct {}

var service ApiImporter

func NewHandler(s ApiImporter) *Handler {
	service = s
	return &Handler{}
}

func (*Handler) ImportData(ctx *iris.Context) {
	data := map[string]interface{}{}
	contentReader := strings.NewReader(string(ctx.PostBody()))
	json.NewDecoder(contentReader).Decode(&data)
	username := data["username"].(string)
	password := data["password"].(string)

	err, auth := service.Authenticate(AuthUrl, username, password)
	if err != nil {
		ctx.JSON(400, map[string]interface{}{
			"code": "NubankAuthenticationError",
			"message": err,
		})
		return
	}

	err, bills := service.GetBillsSummary(auth["url"], auth["token"])
	if err != nil {
		ctx.JSON(400, map[string]interface{}{
			"code": "NubankBillsSummaryError",
			"message": err,
		})
		return
	}

	transactionService := transactions.NewService(transactions.NewDao())

	for _, bill := range bills {
		err, items := service.GetBillItems(bill["link"].(string), auth["token"])
		if err != nil {
			ctx.JSON(400, map[string]interface{}{
				"code": "NubankBillTransactionsError",
				"message": err,
			})
			return
		}
		dueDate, _ := time.Parse("2006-01-02", bill["dueDate"].(string))
		parent := transactions.Transaction{
			Description: "Nubank",
			Amount: bill["paid"].(float64),
			CreatedAt: dueDate,
			Type: "expense",
			Tags: []string{"nubank,creditcard"},
		}
		err, parentId := transactionService.Save(parent)
		if err != nil {
			log.Println("Failed trying to save a transaction.", err)
			continue
		}

		for _, item := range items {
			date, _ := time.Parse("2006-01-02", item["date"].(string))
			children := transactions.Transaction{
				Description: item["title"].(string),
				Amount: item["amount"].(float64),
				CreatedAt: date,
				Type: "expense",
				ParentId: parentId,
			}
			if err, _ := transactionService.Save(children); err != nil {
				log.Println("Failed saving a child transaction with parentId=", parentId, err)
				continue
			}
		}
	}
	
	ctx.JSON(200, "")
}
