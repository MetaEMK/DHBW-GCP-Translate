package config

import (
	"fmt"
	"os"
)

func GetDatabaseConnectionString() (string, error) {
    hostname := os.Getenv("DB_HOSTNAME")
    if hostname == "" {
        hostname = "localhost"
    }

    port := os.Getenv("DB_PORT")
    if port == "" {
        port = "5432"
    }

    username := os.Getenv("DB_USERNAME")
    if username == "" {
        username = "postgres"
    }

    password := os.Getenv("DB_PASSWORD")  
    if password == "" {
        password = "your_password"
        //return "", errors.New("Database password is required")
    }

    dbname := os.Getenv("DB_NAME")
    if dbname == "" {
        dbname = "postgres"  
    }

    connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        hostname, port, username, password, dbname)

    return connectionString, nil
}

