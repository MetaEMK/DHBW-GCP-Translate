package config

import (
	"context"

	"cloud.google.com/go/translate"
	"google.golang.org/api/option"
)

var client *translate.Client

func GetTranslatorClient(ctx context.Context) *translate.Client {
    return client
}

func CreateTranslatorClient(ctx context.Context, keyfilePath *string) error {
    var cl *translate.Client
    var err error

    if keyfilePath == nil {
        cl, err = translate.NewClient(ctx)
    } else {
        cl, err = translate.NewClient(ctx, option.WithCredentialsFile(*keyfilePath))
    }

    client = cl

    return err
}
