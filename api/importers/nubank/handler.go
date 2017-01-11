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

func (*Handler) Hello(ctx *iris.Context) {
	ctx.JSON(200, "")
}

func (*Handler) ImportData(ctx *iris.Context) {
	var data map[string]string
	contentReader := strings.NewReader(string(ctx.PostBody()))
	err := json.NewDecoder(contentReader).Decode(&data)
	if err != nil {
		ctx.JSON(409, map[string]string{
			"code": "InvalidRequestError",
			"message": "Credential invalid.",
		})
		return
	}

	_, discovery := service.Discover()
	err, auth := service.Authenticate(discovery["authUrl"], data["username"], data["password"])
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
		log.Printf("Got %d items from bill id=%s", len(items), bill["id"])

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
			// Disabled fetching details and tags because "too many requests" =/
			/*
			log.Printf("Fetching details about item %s", item["id"])
			url := fmt.Sprintf("%s/%s", TransactionDetailsUrlBase, item["transactionId"])
			err, details := service.GetTransactionDetails(url, auth["token"])
			if err != nil {
				log.Println("Problems to get details for the transsaction.", err)
				continue
			}
			date, _ := details["date"].(time.Time)
			*/
			date, _ := time.Parse("2006-01-02", item["date"].(string))
			children := transactions.Transaction{
				Description: item["title"].(string),
				DueDate: date,
				Type: "expense",
				ParentId: parentId,
				Amount: item["amount"].(float64),
				// Amount: details["amount"].(float64),
				// Tags: details["tags"].([]string),
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
