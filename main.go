package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"os"
)

func main() {
	// Configuration
	listenPort := EnvDefault("PORT", "8080")
	InitLoggers(os.Stdout, os.Stdout, os.Stderr)
	Info.Println("Starting...")

	// Init a new AWS session
	awsSession := CreateAWSSession("eu-west-2")

	// Create the router instance
	router := chi.NewRouter()
	RegisterMiddleware(router)
	RegisterRoutes(router, awsSession)

	// Listen on specified port
	Info.Printf("Server running: http://localhost:%s", listenPort)
	serverErr := http.ListenAndServe(fmt.Sprintf(":%s", listenPort), router)
	if serverErr != nil {
		Error.Fatal(serverErr)
	}
}
