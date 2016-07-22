#!/bin/sh

PUBLIC_DIR="webapp/public"
DEST_DIR="api/webpage"

# index.html -> Html.go
content=$(sed -e 's/`/'"'"'/g' $PUBLIC_DIR/index.html)
read -r -d '' CODE << EOM
package webpage
import "github.com/kataras/iris"
func ServeHtml(ctx *iris.Context) {
	content := \`
		$content
	\`
	ctx.Text(iris.StatusOK, content)
	ctx.SetContentType("text/html")
}
EOM
echo "$CODE" > "$DEST_DIR/html.go"

# app.js -> Js.go
content=$(sed -e 's/`/'"'"'/g' $PUBLIC_DIR/app.js)
read -r -d '' CODE << EOM
package webpage
import "github.com/kataras/iris"
func ServeJs(ctx *iris.Context) {
	content := \`
		$content
	\`
	ctx.Text(iris.StatusOK, content)
	ctx.SetContentType("text/javascript")
}
EOM
printf "%s\n" "$CODE" > "$DEST_DIR/js.go"

# app.css -> Css.go
content=$(sed -e 's/`/'"'"'/g' $PUBLIC_DIR/app.css)
read -r -d '' CODE << EOM
package webpage
import "github.com/kataras/iris"
func ServeCss(ctx *iris.Context) {
	content := \`
		$content
	\`
	ctx.Text(iris.StatusOK, content)
	ctx.SetContentType("text/css")
}
EOM
printf "%s\n" "$CODE" > "$DEST_DIR/css.go"