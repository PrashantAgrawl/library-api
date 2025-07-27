package main

import (
    "log"
    "library-api/config"
    "library-api/internal/database"
    "library-api/internal/router"
)

func main() {
    cfg := config.LoadConfig()

    db, err := database.InitDB(cfg)
    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }
    defer db.Close()

    r := router.SetupRouter(db)
    log.Println("Starting Library API server on port 8080...")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
