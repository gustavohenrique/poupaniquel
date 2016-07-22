package reports

import (
	"log"
	"time"
	"errors"
	"fmt"

	"github.com/kataras/iris"
)

type Handler struct {}

var service Reporter

func NewHandler(s Reporter) *Handler {
	service = s
	return &Handler{}
}

func (*Handler) Report(ctx *iris.Context) {
	err, params := getParamsFrom(ctx)
	var result interface{}
	status := 200
	if err == nil {
		err, result = service.ByTag(params)
	}
	if err != nil {
		status = 400
		result = map[string]interface{}{
			"code": "ReportError",
			"message": fmt.Sprintf("%s", err),
		}
		log.Println("Error generating report by tag.", err)
	}
	ctx.JSON(status, result)
}

func getParamsFrom(ctx *iris.Context) (err error, params map[string]interface{}) {
	tag := ctx.URLParam("tag")
	transactionType := ctx.URLParam("type")
	from := ctx.URLParam("from")
	until := ctx.URLParam("until")
	if tag == "" || from == "" || until == "" || transactionType == "" {
		err = errors.New("The url params 'tag', 'type', 'from' and 'until' are required.")
		return err, params
	}
	startDate, err := time.Parse("2006-01-02", from)
	endDate, err := time.Parse("2006-01-02", until)
	if err != nil {
		return err, params
	}
	return err, map[string]interface{} {
		"tag": tag,
		"type": transactionType,
		"startDate": startDate,
		"endDate": endDate,
	}
}