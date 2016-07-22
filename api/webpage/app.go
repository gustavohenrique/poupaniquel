package webpage

import "github.com/kataras/iris"

func New(server *iris.Framework) {
	server.Get("/", ServeHtml)
	server.Get("/app.js", ServeJs)
	server.Get("/app.css", ServeCss)
	server.Get("/docs", func (ctx *iris.Context) {
		ctx.Redirect("http://docs.poupaniquel.apiary.io/", 301)
	})
}
