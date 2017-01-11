package webpage
import "gopkg.in/gin-gonic/gin.v1"
func ServeHtml(ctx *gin.Context) {
	content := `
		<!DOCTYPE html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=yes">
  <title>Poupaniquel</title>
  <link rel="stylesheet" href="http://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.6.3/css/font-awesome.min.css">
  <link rel="stylesheet" type="text/css" href="app.css">
  <script src="app.js"></script>
  <script>require('initialize');</script>
</head>
<body>
  <div id="app"></div>
  <script src="http://cdnjs.cloudflare.com/ajax/libs/fastclick/1.0.6/fastclick.min.js"></script>
  <script type="text/javascript">
    if ('addEventListener' in document) {
      document.addEventListener('DOMContentLoaded', function() {
        FastClick.attach(document.body);
      }, false);
    }
  </script>
</body>
</html>
	`
  ctx.Writer.Header().Set("Content-Type", "text/html")
  ctx.String(200, content)
}
