package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/metaemk/dhbw-gcp-translate/config"
)

var dbConn *sql.DB
var confPath string

func getDatabase(ctx context.Context) (*sql.Conn, error) {
    if dbConn == nil {
        InitDatabase(confPath)
    }
    conn, err := dbConn.Conn(ctx)
    if err != nil {
        return nil, err
    }

    err = conn.PingContext(ctx)

    return conn, err
}

func InitDatabase(configPath string) error {
   confPath = configPath
   c, err := config.GetDatabaseConfig(configPath)
   if err != nil {
       return err
   }

   dsn := fmt.Sprintf("user=%s password=%s database=%s", c.DBUsername, c.DBPassword, c.DBName)
   config, err := pgx.ParseConfig(dsn)
   if err != nil {
       return err
   }

   var opts []cloudsqlconn.Option
   if c.DBHostname != "" {
       opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
   }
   d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
   if err != nil {
       return err
   }

   config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
       return d.Dial(ctx, c.DBInstanceConnectionName)
   }
   dbURI := stdlib.RegisterConnConfig(config)
   dbPool, err := sql.Open("pgx", dbURI)
   if err != nil {
       return fmt.Errorf("sql.Open: %w", err)
   }

   dbConn = dbPool

   err = createTable()

   return err
}

func createTable() error {
    println("Creating table")
    ctx := context.Background()
    conn, err := getDatabase(ctx)
    if err != nil {
        return err
    }
    defer conn.Close()

    query := `
        CREATE TABLE IF NOT EXISTS translation (
            id SERIAL,
            insert_time DATE,
            promt_hash TEXT,
            target_lang TEXT,
            translation_text TEXT
        );
    `
    _, err = conn.ExecContext(ctx, query)
    println("Table created")

    return err
}
