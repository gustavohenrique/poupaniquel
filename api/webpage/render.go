package webpage

import "github.com/kataras/iris"

func NewController() {
	iris.Get("/", Html)
	iris.Get("/app.js", Js)
	iris.Get("/app.css", Css)
}

// func Css(ctx *iris.Context) {
// 	ctx.Text(iris.StatusOK, ReadFile("static/app.css"))	
// 	ctx.SetContentType("text/css")
// }
// func ReadFile(path string) string {
// 	content, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(content)
// }