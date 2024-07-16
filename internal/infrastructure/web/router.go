package web

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"shiharaikun/internal/adapter/rest/handler"
	"shiharaikun/internal/usecase"
)

func RegisterRoutes(invoiceUseCase usecase.InvoiceUseCase) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	}).Methods(http.MethodGet)

	invoiceRouter := r.PathPrefix("/api/invoices").Subrouter()
	invoiceHandler := handler.NewInvoiceHandler(invoiceUseCase)
	invoiceRouter.HandleFunc("", invoiceHandler.CreateInvoiceHandler).Methods(http.MethodPost)
	invoiceRouter.HandleFunc("", invoiceHandler.ListInvoicesByDueDateHandler).Methods(http.MethodGet)

	return r
}
