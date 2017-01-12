package reports_test

import (
	"testing"
	"encoding/json"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/appleboy/gofight.v2"

	"github.com/gustavohenrique/poupaniquel/api"
	"github.com/gustavohenrique/poupaniquel/api/reports"
	"github.com/gustavohenrique/poupaniquel/api/reports/fake"
)

func TestReportsSpec(t *testing.T) {

	Convey("/reports", t, func() {
		server :=  api.NewServer()
		client := gofight.New()

		reports.New(server, map[string]interface{}{
			"baseUrl": "/api/v1",
			"service": &fake.Service{},
		})
		
		Convey("Should returns an error if no URL params was sent", func() {
			client.GET("/api/v1/reports").
				SetDebug(true).
    			Run(server, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
    				So(r.Code, ShouldEqual, 409)

    				var resp map[string]string
    				raw := []byte(r.Body.String())

    				json.Unmarshal(raw, &resp)
					So(resp, ShouldContainKey, "code")
					So(resp, ShouldContainKey, "message")
    			})

		})
		
		Convey("Should returns data for chart's usage", func() {
			client.GET("/api/v1/reports").
				SetQuery(gofight.H{
			      "from": "2016-01-01",
			      "until": "2016-12-01",
			      "type": "expense",
			      "tag": "creditcard",
			    }).
			    Run(server, func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
    				So(r.Code, ShouldEqual, 200)

    				var resp []map[string]interface{}
    				raw := []byte(r.Body.String())
    				json.Unmarshal(raw, &resp)

					first := resp[1]
					month := first["month"].(string)
					amount := first["amount"].(float64)
					total := first["total"].(float64)
					So(month, ShouldEqual, "2016-01")
					So(amount, ShouldEqual, 50.2)
					So(total, ShouldEqual, 100)

					second := resp[2]
					month = second["month"].(string)
					amount = second["amount"].(float64)
					total = second["total"].(float64)
					So(month, ShouldEqual, "2016-02")
					So(amount, ShouldEqual, 320.2)
					So(total, ShouldEqual, 890)
    			})

		})

	})
}
