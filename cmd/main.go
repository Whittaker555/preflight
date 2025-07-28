package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/whittaker555/preflight/internal/logger"
	"github.com/whittaker555/preflight/internal/routes"
)

type runner interface {
	Run(addr ...string) error
}

func runServer(r runner, port string) error {
	logger.Log.Infof("PreFlight API running on port %s", port)
	return r.Run(":" + port)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using defaults")
	}

	logger.Init()

	r := gin.Default()
	routes.RegisterRoutes(r)

	// If running inside AWS Lambda, use the Lambda handler
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		ginLambda := ginadapter.NewV2(r)
		lambda.Start(ginLambda.ProxyWithContext)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := runServer(r, port); err != nil {
		logger.Log.Fatalf("server error: %v", err)
	}
}
