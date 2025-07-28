package main

import (
    "log"
    "os"

    ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"

    "github.com/whittaker555/preflight/internal/logger"
    "github.com/whittaker555/preflight/internal/routes"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using defaults")
    }

    logger.Init()

    r := gin.Default()
    routes.RegisterRoutes(r)

    // If running inside AWS Lambda, use the Lambda handler
    if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
        ginLambda := ginadapter.New(r)
        lambda.Start(ginLambda.Proxy)
        return
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    logger.Log.Infof("PreFlight API running on port %s", port)
    r.Run(":" + port)
}
