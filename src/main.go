package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/metaemk/dhbw-gcp-translate/api"
	"github.com/metaemk/dhbw-gcp-translate/config"
	"github.com/metaemk/dhbw-gcp-translate/database"
)

func main() {
    var configPath string
    flag.StringVar(&configPath, "c", "/etc/translator/config.yaml", "path of the config file")
    flag.Parse()

    ctx := context.Background()

    fmt.Println("Trying to connect to the translator")
    err := config.CreateTranslatorClient(ctx, configPath)
    if err != nil {
        fmt.Println("Could not connect to the translator")
        panic(err.Error())
    }
    fmt.Println("Translator connection established")

    fmt.Println("Trying to connect to the database")
    for {
        err = database.InitDatabase(configPath)
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
