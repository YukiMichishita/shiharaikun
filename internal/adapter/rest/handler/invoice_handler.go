package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"shiharaikun/internal/usecase"
	"shiharaikun/internal/usecase/model"
	"strconv"
	"time"
)

type InvoiceHandler struct {
	useCase usecase.InvoiceUseCase
}

func NewInvoiceHandler(useCase usecase.InvoiceUseCase) *InvoiceHandler {
	return &InvoiceHandler{useCase: useCase}
}

func (i *InvoiceHandler) CreateInvoiceHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req model.CreateInvoiceRequest
	if err := json.Unmarshal(body, &req); err != nil {
		log.Print(err)
		http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
		return
	}
	tenantID := r.Header.Get("x-Tenant-ID")
	req.TenantID, err = strconv.Atoi(tenantID)
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid tenant ID", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	resp, err := i.useCase.CreateInvoice(ctx, &req)
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to create invoice", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Print(err)
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
		return
	}
}

func (i *InvoiceHandler) ListInvoicesByDueDateHandler(w http.ResponseWriter, r *http.Request) {
	tenantID := r.Header.Get("x-Tenant-ID")
	tenantIDInt, err := strconv.Atoi(tenantID)
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid tenant ID", http.StatusBadRequest)
		return
	}

	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	if startDate == "" || endDate == "" {
		log.Print(err)
		http.Error(w, "Invalid date range", http.StatusBadRequest)
		return
	}

	sd, err := time.Parse(time.DateOnly, startDate)
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid start date", http.StatusBadRequest)
		return
	}
	ed, err := time.Parse(time.DateOnly, endDate)
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid end date", http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	req := &model.GetInvoicesRequest{
		TenantID:  tenantIDInt,
		StartDate: sd,
		EndDate:   ed,
	}
	resp, err := i.useCase.GetInvoices(ctx, req)
	if err != nil {
		log.Print(err)
		http.Error(w, "Unable to get invoices", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Print(err)
		http.Error(w, "Unable to encode response", http.StatusInternalServerError)
		return
	}
}
