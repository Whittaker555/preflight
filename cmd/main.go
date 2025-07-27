package main

import (
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
    "os"
    "github.com/yourusername/preflight/internal/routes"
    "github.com/yourusername/preflight/internal/logger"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using defaults")
    }

    logger.Init()

    r := gin.Default()
    routes.RegisterRoutes(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    logger.Log.Infof("PreFlight API running on port %s", port)
    r.Run(":" + port)
}
