package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/metaemk/dhbw-gcp-translate/config"
)

var templatePath string

func CreateServer(config *config.HttpConfig) *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(cors.Default())

    basePath := config.TemplatesPath
    templatePath = basePath

    r.LoadHTMLFiles(basePath + "/index.html", basePath + "/styles.css")
    r.GET("/", getWebsiteHtml)
    r.GET("/styles.css", getWebsiteCss)
    r.GET("/error-handler.js", getWebsiteJs)

    apiGroup := r.Group("/api")
    apiGroup.POST("/translate", translateHandler)

    return r
}

func getWebsiteHtml(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "Main website",
    })
}

func getWebsiteCss(c *gin.Context) {
    c.File(templatePath + "/styles.css")
}

func getWebsiteJs(c *gin.Context) {
    c.File(templatePath + "/error-handler.js")
}
