package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"todo-service/pkg/handler"
	"todo-service/pkg/usecase/task/repository"
	"todo-service/pkg/usecase/task/service"
)

func main() {
	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username: "root",
		Password: "example",
	})

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	// Initialize repository, service, and handler
	todoRepo := repository.NewMongoRepository(client)
	todoService := service.NewService(todoRepo)
	todoHandler := handler.NewHandler(todoService)

	// Set up routes
	http.HandleFunc("/api/v1/add", todoHandler.AddTask)
	http.HandleFunc("/api/v1/complite", todoHandler.CompleteTask)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
