package reports

import (
	"github.com/kataras/iris"
	"log"
	"time"
	"strings"
	"errors"
	"fmt"
)

type Controller struct {}

var service = NewService()

func NewController(base string) {
	api := &Controller{}
	iris.Get(base + "/reports", api.Report)
}

func (this *Controller) Report(ctx *iris.Context) {
	err, params := getParamsFrom(ctx)
	var result interface{}
	status := 200
	if err == nil {
		err, result = service.ByTags(params)
	}
	if err != nil {
		status = 400
		result = map[string]interface{}{
			"code": "ReportError",
			"message": fmt.Sprintf("%s", err),
		}
		log.Println("Error generating report by tags.", err)
	}
	ctx.JSON(status, result)
}

func getParamsFrom(ctx *iris.Context) (err error, params map[string]interface{}) {
	list := ctx.URLParam("tags")
	transactionType := ctx.URLParam("type")
	from := ctx.URLParam("from")
	until := ctx.URLParam("until")
	if list == "" || from == "" || until == "" || transactionType == "" {
		err = errors.New("The url params 'tags', 'type', 'from' and 'until' are required.")
		return err, params
	}
	startDate, err := time.Parse("2006-01-02", from)
	endDate, err := time.Parse("2006-01-02", until)
	if err != nil {
		return err, params
	}
	return err, map[string]interface{} {
		"tags": strings.Split(list, ","),
		"type": transactionType,
		"startDate": startDate,
		"endDate": endDate,
	}
}