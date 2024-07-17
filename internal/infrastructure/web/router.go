package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"shiharaikun/internal/adapter/rest/handler"
	"shiharaikun/internal/adapter/rest/middleware"
	"shiharaikun/internal/usecase"
)

func RegisterRoutes(userUseCase usecase.UserUseCase, invoiceUseCase usecase.InvoiceUseCase) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	}).Methods(http.MethodGet)

	authMiddleware := middleware.NewAuthMiddleware(userUseCase)
	r.Use(authMiddleware.HandleSession)

	invoiceRouter := r.PathPrefix("/api/invoices").Subrouter()
	invoiceHandler := handler.NewInvoiceHandler(invoiceUseCase)
	invoiceRouter.HandleFunc("", invoiceHandler.CreateInvoiceHandler).Methods(http.MethodPost)
	invoiceRouter.HandleFunc("", invoiceHandler.ListInvoicesByDueDateHandler).Methods(http.MethodGet)

	return r
}
