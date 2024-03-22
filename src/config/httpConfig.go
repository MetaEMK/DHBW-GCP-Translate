package config

import (
	"os"

	"gopkg.in/yaml.v2"
)


type HttpConfig struct {
    // Port of the HTTP server
    Port int `yaml:"port"`

    // StaticDir is the directory where the static files are stored
    TemplatesPath string `yaml:"templates_path"`
}

func GetHttpConfig(configPath string) (HttpConfig, error) {
    if configPath == "" {
        configPath = "./config.yaml"
    }

    yamlFile, err := os.Open(configPath)
    if err != nil {
        return HttpConfig{}, err
    }
    defer yamlFile.Close()

    var config HttpConfig

    err = yaml.NewDecoder(yamlFile).Decode(&config)
    if err != nil {
        return HttpConfig{}, err
    }

    return config, nil
}
