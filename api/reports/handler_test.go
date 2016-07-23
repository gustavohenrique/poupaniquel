package reports_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/kataras/iris"

	"github.com/gustavohenrique/poupaniquel/api"
	"github.com/gustavohenrique/poupaniquel/api/reports"
	"github.com/gustavohenrique/poupaniquel/api/reports/fake"
)

func TestReportsSpec(t *testing.T) {

	Convey("/reports", t, func() {
		server :=  api.NewServer()
		reports.New(server, map[string]interface{}{
			"baseUrl": "/api/v1",
			"service": &fake.Service{},
		})
		e := iris.NewTester(server, t)
		
		Convey("Should returns an error if no URL params was sent", func() {
			resp := e.GET("/reports").Expect()
			So(resp.Raw().StatusCode, ShouldEqual, 409)

			raw := resp.JSON().Object().Raw()
			So(raw, ShouldContainKey, "code")
			So(raw, ShouldContainKey, "message")
		})
		
		Convey("Should returns data for chart's usage", func() {
			resp := e.GET("/reports").
					  WithQuery("from", "2016-01-01").
					  WithQuery("until", "2016-12-01").
					  WithQuery("type", "expense").
					  WithQuery("tag", "creditcard").
					  Expect()

			So(resp.Raw().StatusCode, ShouldEqual, 200)

			raw := resp.JSON().Array().Raw()
			first := raw[1].(map[string]interface{})
			month := first["month"].(string)
			amount := first["amount"].(float64)
			total := first["total"].(float64)
			So(month, ShouldEqual, "2016-01")
			So(amount, ShouldEqual, 50.2)
			So(total, ShouldEqual, 100)

			second := raw[2].(map[string]interface{})
			month = second["month"].(string)
			amount = second["amount"].(float64)
			total = second["total"].(float64)
			So(month, ShouldEqual, "2016-02")
			So(amount, ShouldEqual, 320.2)
			So(total, ShouldEqual, 890)
		})
	})
}

/*
func IrisTester(t *testing.T) *httpexpect.Expect {
	server :=  api.NewServer()
	reports.NewController(server, "/api/v1", nil)
	// service := &ServiceFake{}
	// reports.NewController(server, "/api/v1", service)
	handler := server.ListenVirtual().Handler

	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: "/api/v1",
		Client: &http.Client{
			Transport: httpexpect.NewFastBinder(handler),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			// httpexpect.NewDebugPrinter(t, true),
		},
	})
}
*/