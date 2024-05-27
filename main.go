package main

import (
	"StudentRegistryAPI/router"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get the port number from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("PORT environment variable is not set")
	}

	r := router.NewRouter()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :%s: %v\n", port, err)
		}
	}()

	// Print the server running message
	fmt.Printf("server running on port %s\n", port)

	// Block main goroutine
	select {}
}
