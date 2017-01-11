package webpage

import "gopkg.in/gin-gonic/gin.v1"

func New(server *gin.Engine) {
	server.GET("/", ServeHtml)
	server.GET("/app.js", ServeJs)
	server.GET("/app.css", ServeCss)
	server.GET("/docs", func (ctx *gin.Context) {
		ctx.Redirect(301, "http://docs.poupaniquel.apiary.io/")
	})
}
