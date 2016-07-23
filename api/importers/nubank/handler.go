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
	log.Println("Authenticated in Nubank.")

	err, bills := service.GetBillsSummary(auth["url"], auth["token"])
	if err != nil {
		ctx.JSON(400, map[string]interface{}{
			"code": "NubankBillsSummaryError",
			"message": err,
		})
		return
	}
	log.Printf("Got %d bills from Nubank.", len(bills))

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
		log.Printf("Got %d items from bill id = %s", len(items), bill["id"])

		dueDate, _ := time.Parse("2006-01-02", bill["dueDate"].(string))
		parent := transactions.Transaction{
			Description: "Nubank",
			Amount: bill["paid"].(float64),
			DueDate: dueDate,
			Type: "expense",
			Tags: []string{"|nubank|,|creditcard|"},
		}
		err, parentId := transactionService.Save(parent)
		if err != nil {
			log.Println("Failed trying to save a transaction.", err)
			continue
		}
		log.Printf("Transaction #%d was saved.", parentId)

		for _, item := range items {
			date, _ := time.Parse("2006-01-02", item["date"].(string))
			children := transactions.Transaction{
				Description: item["title"].(string),
				Amount: item["amount"].(float64),
				DueDate: date,
				Type: "expense",
				ParentId: parentId,
			}
			if err, _ := transactionService.Save(children); err != nil {
				log.Println("Failed saving a child transaction with parentId=", parentId, err)
				continue
			}
		}
		log.Printf("Saved %d children for #%d\n\n", len(items), parentId)
	}

	log.Println("Nubank import is finished.")
	
	ctx.JSON(200, "")
}