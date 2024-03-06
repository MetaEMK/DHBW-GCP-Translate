package api

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/metaemk/dhbw-gcp-translate/config"
	"golang.org/x/text/language"
)

type TranslateRequest struct {
    Text string `json:"text"`
    TargetLanguage language.Tag `json:"targetLanguage"`
}

func translateHandler(c *gin.Context) {
    var err error
    ctx := context.Background()
    client, err := config.NewTranslatorClient()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    defer client.Close()

    var request TranslateRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    translation, err := client.Translate(ctx, []string{request.Text}, request.TargetLanguage, nil)

    if err != nil {
        println(err)
        c.JSON(500, gin.H{"error": err.Error()})
    } else {
        c.String(200, translation[0].Text)
    }
}

func getSupportedLanguagesHandler(c *gin.Context) {
    var err error
    ctx := context.Background()
    client, err := config.NewTranslatorClient()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    defer client.Close()

    var targetLanguage language.Tag
    // DEBUG ONLY
    targetLanguage = language.German
    languages, err := client.SupportedLanguages(ctx, targetLanguage)

    if err != nil {
        c.JSON(500, gin.H{"error": err})
    } else {
        c.JSON(200, gin.H{"message": languages})
    }
}

func detectLanguageHandler(c *gin.Context) {
    var err error
    ctx := context.Background()

    client, err := config.NewTranslatorClient()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    defer client.Close()

    text := "Hello World"
    langs, err := client.DetectLanguage(ctx, []string{text})

    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
    } else {
        c.JSON(200, gin.H{"message": langs})
    }
} 
