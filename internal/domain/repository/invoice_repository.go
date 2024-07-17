package repository

import (
	"context"
	"shiharaikun/internal/domain/model"
	"time"
)

type InvoiceRepository interface {
	CreateInvoice(ctx context.Context, input *model.Invoice) (*model.Invoice, error)
	ListInvoicesByDueDate(ctx context.Context, input *ListInvoiceByDueDateInput) ([]*model.Invoice, error)
}

type ListInvoiceByDueDateInput struct {
	StartDate time.Time
	EndDate   time.Time
	CompanyID int32
}
