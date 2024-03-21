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

    result, entryFound, err := database.GetTranslation(ctx, request)
    if err != nil || entryFound{
        if entryFound {
            println("entry found")
        }
        return
    }

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

    go database.SaveTranslation(ctx, result)
    return
}
