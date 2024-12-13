package main

import (
	"log"
	"net/http"
	db "refactoring_tech/data"
	"refactoring_tech/domain/service"
	"refactoring_tech/presentation/handlers"
)

func main() {
	// init config
	config := db.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "1234",
		DbName:   "refactoring_tech",
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

	// init handlers
	techHandler := handlers.NewTechHandler(controller)
	producerHandler := handlers.NewProducerHandler(controller)

	startListening(techHandler, producerHandler)
}

func startListening(techHandler *handlers.TechHandler, producerHandler *handlers.ProducerHandler) {
	http.HandleFunc("/techs", techHandler.GetTechsHandler)
	http.HandleFunc("/tech", techHandler.TechHandler)
	http.HandleFunc("/producer", producerHandler.ProducerHandler)

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
