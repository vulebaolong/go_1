package swagger

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/swaggest/openapi-go/openapi3"
)

func Start(ginEngine *gin.Engine) {

	ginEngine.GET("docs", func(ctx *gin.Context) {
		html := `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <meta name="color-scheme" content="light dark" />
  <title>MegaPro Chat Bot Swagger</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.32.2/swagger-ui.css" />
</head>
<body style="background: #f8f9fa;">
  <div id="swagger-ui"></div>

  <script src="https://unpkg.com/swagger-ui-dist@5.32.2/swagger-ui-bundle.js"></script>
  <script src="https://unpkg.com/swagger-ui-dist@5.32.2/swagger-ui-standalone-preset.js"></script>

  <script>
    window.onload = function () {
      window.ui = SwaggerUIBundle({
        url: "/docs.json",
        dom_id: "#swagger-ui",
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis,
          SwaggerUIStandalonePreset
        ],
        layout: "StandaloneLayout"
      });
    };
  </script>
</body>
</html>
`

		ctx.Data(200, "text/html charset=utf-8", []byte(html))
	})

	ginEngine.GET("/docs.json", func(ctx *gin.Context) {
		reflector := &openapi3.Reflector{}

		reflector.Spec = &openapi3.Spec{Openapi: "3.0.3"}

		reflector.Spec.Info.
			WithTitle("Golang Backend").
			WithVersion("1.0.0").
			WithDescription("Description")

		docJson, err := json.MarshalIndent(reflector.Spec, "", "     ")
		if err != nil {
			ctx.JSON(400, "Lỗi")
		}

		ctx.Data(200, "application/json charset=utf-8", docJson)
	})

}
