package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
    DBHostname string `yaml:"db_hostname"`
    DBPort     string `yaml:"db_port"`
    DBUsername string `yaml:"db_username"`
    DBPassword string `yaml:"db_password"`
    DBName     string `yaml:"db_name"`
    DBInstanceConnectionName string `yaml:"db_instance_connection_name"`
}

func GetDatabaseConfig(configPath string) (DatabaseConfig, error) {
    if configPath == "" {
        configPath = "./config.yaml"
    }

    yamlFile, err := os.Open(configPath)
    if err != nil {
        return DatabaseConfig{}, err
    }
    defer yamlFile.Close()

    var config DatabaseConfig

    err = yaml.NewDecoder(yamlFile).Decode(&config)
    if err != nil {
        return DatabaseConfig{}, err
    }

    return config, nil
}
