package translator

import (
	"context"
	"errors"

	"github.com/metaemk/dhbw-gcp-translate/config"
	"github.com/metaemk/dhbw-gcp-translate/database"
	"github.com/metaemk/dhbw-gcp-translate/model"
)


func Translate(ctx context.Context, request model.TranslateRequest) (result model.TranslationResult, err error) {
    client := config.GetTranslatorClient(ctx)

    // Check if the translation is already in cache
    result, entryFound, err := database.GetTranslation(ctx, request)
    if err != nil || entryFound {
        return
    }

    // Translate the text
    translation, err := client.Translate(ctx, []string{request.Text}, request.TargetLanguage, nil)
    if err != nil {
        return
    }

    if len(translation) == 1 {
        result.Translation = translation[0].Text
        result.SourceLanguage = translation[0].Source
        result.TargetLanguage = request.TargetLanguage
        result.Text = request.Text
    } else {
        err = errors.New("client responded with more than 1 translation or 0")
        println("client responded with more than 1 translation")
    }

    // Save the translation to the database
    go database.SaveTranslation(ctx, result)
    return
}
