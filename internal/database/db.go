package database

import (
    "database/sql"
    "fmt"
    "library-api/config"
    "log"

    _ "github.com/lib/pq"
)

func InitDB(cfg config.DBConfig) (*sql.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Connected to the database successfully.")
    return db, nil
}
