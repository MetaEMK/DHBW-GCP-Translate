package database

import (
	"context"
	"time"

	"github.com/metaemk/dhbw-gcp-translate/model"
	"golang.org/x/text/language"
)

func GetTranslation(ctx context.Context, request model.TranslateRequest) (translation model.TranslationResult, entryFound bool, err error) {
    entryFound = false
    dbConn, err := getDatabase(ctx)
    if err != nil {
        return 
    }
    defer dbConn.Close()

    hash := request.HashText()

    result, err := dbConn.QueryContext(
        ctx,
        "SELECT promt_hash, target_lang, translation_text FROM translation WHERE promt_hash = $1 AND target_lang = $2 LIMIT 1",
        hash,
        request.TargetLanguage.String(),
    )
    if err != nil {
        return 
    }
    defer result.Close()

    for result.Next() {
        var test string
        var target_lang string
        var translation_text string

        err = result.Scan(&test, &target_lang, &translation_text)
        if err != nil {
            return 
        }

        translation.Translation = translation_text

        var t language.Tag
        t, err = language.Parse(target_lang)
        if err != nil {
            return 
        }

        translation.TargetLanguage = t

        entryFound = true
        break
    }

    return
}

func SaveTranslation(ctx context.Context, translation model.TranslationResult) {
    dbConn, err := getDatabase(ctx)
    if err != nil {
        return
    }
    defer dbConn.Close()

    hash := translation.HashText()
    insertTime := time.Now().UTC()

    _, err = dbConn.ExecContext(
        ctx,
        "INSERT INTO translation (insert_time, promt_hash, target_lang, translation_text) VALUES ($1, $2 , $3, $4)",
        insertTime,
        hash,
        translation.TargetLanguage.String(),
        translation.Translation,
    )

    if err != nil {
        println("Error while inserting translation: %s", err.Error())
    }
}
