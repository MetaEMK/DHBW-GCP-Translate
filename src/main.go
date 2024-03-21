package main

import (
	"context"
	"fmt"
	"time"

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

    fmt.Println("Trying to connect to the database")
    for {
        err = database.InitDatabase("")
        if err != nil {
            println(err.Error())
            fmt.Println("Could not connect to the database - retrying in 2 seconds")
            time.Sleep(2 * time.Second)
        } else {
            break
        }
    }

    fmt.Println("Database connection established")

    r := api.CreateServer()
    r.Run(":8080")
}
