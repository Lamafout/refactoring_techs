package main

import (
	"log"
	"net/http"
	db "refactoring_tech/data"
	"refactoring_tech/domain/service"
	"refactoring_tech/presentation"
)

func main() {
	// init config
	config := db.Config{
		Host: "localhost",
		Port: "5432",
		User: "postgres",
		Password: "1234",
		DbName: "refactoring_tech",
	}

	//init connection
	connect, err := db.ConnectToPostgres(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer connect.Close()

	// init repository
	repo := db.NewRepositoryImpl(connect)

	// init use cases
	useCases := service.NewUseCases(repo)

	// init controller
	controller := service.NewControllerImpl(useCases)

	// init handler
	handler := presentation.NewHandler(controller, controller)

	startListening(handler)
}

func startListening(handler *presentation.Handler) {
	http.HandleFunc("/techs", handler.GetTechsHandler)
	http.HandleFunc("/tech", handler.GetConcreteTechHandler)

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}