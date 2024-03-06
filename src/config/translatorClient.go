package config

import (
	"context"

	"cloud.google.com/go/translate"
)


func NewTranslatorClient() (*translate.Client, error) {
    ctx := context.Background()
    client, err := translate.NewClient(ctx)
    if err != nil {
        println(err.Error())
    }

    return client, err
}
