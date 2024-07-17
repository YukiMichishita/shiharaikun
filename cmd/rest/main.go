package main

import (
	"log"
	"net/http"
	"shiharaikun/internal/adapter/repository"
	"shiharaikun/internal/infrastructure/db"
	"shiharaikun/internal/infrastructure/web"
	"shiharaikun/internal/usecase/interactor"
)

func main() {
	invoiceRepo := repository.NewInvoiceRepository()
	userRepo := repository.NewUserRepository()
	invoiceUseCase := interactor.NewInvoiceInterActor(invoiceRepo)
	userUseCase := interactor.NewUserInterActor(userRepo)

	err := db.SetupGen()
	if err != nil {
		log.Println("Error setting up database:", err)
		return
	}

	r := web.RegisterRoutes(userUseCase, invoiceUseCase)
	log.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println("Error starting server:", err)
	}
}
