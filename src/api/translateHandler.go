package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/metaemk/dhbw-gcp-translate/model"
	"github.com/metaemk/dhbw-gcp-translate/translator"
)

func translateHandler(c *gin.Context) {
    var err error
    ctx := context.Background()

    var request model.TranslateRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    translation, err := translator.Translate(ctx, request)

    if err != nil {
        println(err.Error())
        c.JSON(500, gin.H{"error": err.Error()})
    } else {
        c.String(200, translation.Translation)
    }
}

