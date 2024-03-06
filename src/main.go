package main

import "github.com/metaemk/dhbw-gcp-translate/api"

func main() {
    r := api.CreateServer()
    r.Run(":8080")
}
