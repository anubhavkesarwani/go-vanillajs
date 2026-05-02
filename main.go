package main

import (
	"log"
	"net/http"
	"os"
	"frontendmasters.com/reelingit/handlers"
	"frontendmasters.com/reelingit/logger"
	"github.com/joho/godotenv"
)

func main() {
	logInstance := initializeLogger()

	if err := godotenv.Load(); err != nil {
		logInstance.Error("Error loading .env file", err)
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		logInstance.Info("DATABASE_URL is not set")
	}

	movieHandler := handlers.MovieHandler{}
	
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	logInstance.Info("Server starting on " + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server failed to start", err)
		log.Fatalf("Server failed: %v", err)
	}
}

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie-service.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	return logInstance
}
