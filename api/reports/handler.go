package reports

import (
	"log"
	"time"
	"errors"

	"gopkg.in/gin-gonic/gin.v1"
)

type Handler struct {}

var service Reporter

func NewHandler(s Reporter) *Handler {
	service = s
	return &Handler{}
}

func (*Handler) Report(ctx *gin.Context) {
	err, params := getParamsFrom(ctx)
	if err != nil {
		log.Println("Error getting query data to generate report.", err)
		ctx.JSON(409, map[string]interface{}{
			"code": "InsuficientParametersError",
			"message": err,
		})
		return
	}

	err, report := service.ByTag(params)
	if err != nil {
		log.Println("Error generating report by tag.", err)
		ctx.JSON(400, map[string]interface{}{
			"code": "ReportError",
			"message": err,
		})
		return
	}
	ctx.JSON(200, report)
}

func getParamsFrom(ctx *gin.Context) (err error, params map[string]interface{}) {
	tag := ctx.Param("tag")
	transactionType := ctx.Param("type")
	from := ctx.Param("from")
	until := ctx.Param("until")
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