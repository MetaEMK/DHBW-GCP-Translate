package config

import (
	"context"
	"os"

	"cloud.google.com/go/translate"
	"google.golang.org/api/option"
	"gopkg.in/yaml.v2"
)

type TranslatorConfig struct {
    KeyfilePath string `yaml:"keyfile_path"`
}

var client *translate.Client

func GetTranslatorClient(ctx context.Context) *translate.Client {
    return client
}

func CreateTranslatorClient(ctx context.Context, configPath string) error {
    if configPath == "" {
        configPath = "./config.yaml"
    }

    yamlFile, err := os.Open(configPath)
    if err != nil {
        return err
    }
    defer yamlFile.Close()

    var config TranslatorConfig

    err = yaml.NewDecoder(yamlFile).Decode(&config)
    if err != nil {
        return err
    }


    var cl *translate.Client
    if config.KeyfilePath == "" {
        cl, err = translate.NewClient(ctx)
    } else {
        cl, err = translate.NewClient(ctx, option.WithCredentialsFile(config.KeyfilePath))
    }

    client = cl

    return err
}
