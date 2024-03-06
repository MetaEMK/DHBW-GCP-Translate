package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    r.Use(cors.Default())

    r.LoadHTMLFiles("templates/index.html", "templates/styles.css")
    r.GET("/", getWebsiteHtml)
    r.GET("/styles.css", getWebsiteCss)

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
    c.File("templates/styles.css")
}
