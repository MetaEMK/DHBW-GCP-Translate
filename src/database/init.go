package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/metaemk/dhbw-gcp-translate/config"
    _ "github.com/lib/pq"
)

var dbConn *sql.DB

func getDatabase(ctx context.Context) (*sql.Conn, error) {
    if dbConn == nil {
        InitDatabase()
    }
    conn, err := dbConn.Conn(ctx)
    if err != nil {
        return nil, err
    }

    return conn, nil
}

func InitDatabase() error {
    connString, err := config.GetDatabaseConnectionString()
    if err != nil {
        return err
    }

    conn, err := sql.Open("postgres", connString)
    if err != nil {
        return err
    }

    dbConn = conn

    err = createTable()
    return err 
}

func createTable() error {
    ctx := context.Background()
    conn, err:= getDatabase(ctx)
    if err != nil {
        return errors.New("Could not setup the database")
    }
    defer conn.Close()

    query := `
        CREATE TABLE IF NOT EXISTS translation (
            id SERIAL,
            insert_time DATE,
            promt_hash TEXT,
            target_lang TEXT,
            translation_text TEXT
        )
    `
    _, err = conn.ExecContext(ctx, query)

    return err
}
