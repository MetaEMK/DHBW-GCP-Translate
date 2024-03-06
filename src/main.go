package main

import (
	"context"

	"github.com/metaemk/dhbw-gcp-translate/api"
	"github.com/metaemk/dhbw-gcp-translate/config"
	"github.com/metaemk/dhbw-gcp-translate/database"
)

func main() {
    ctx := context.Background()
    str := "/home/jan/Downloads/cc2-translator-0aeb3bb8fb55.json"
    err := config.CreateTranslatorClient(ctx, &str)
    if err != nil {
        panic(err.Error())
    }

    err = database.InitDatabase()
    if err != nil {
        panic(err.Error())
    }

    r := api.CreateServer()
    r.Run(":8080")
}
