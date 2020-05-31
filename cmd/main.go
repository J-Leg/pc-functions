package main

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	f "github.com/J-Leg/pc-functions"
	"log"
	"os"
)

func main() {
	funcframework.RegisterHTTPFunction("/monthly", f.ProcessMonthly)
	funcframework.RegisterHTTPFunction("/daily", f.ProcessDaily)
	funcframework.RegisterHTTPFunction("/recover", f.Recover)

	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
