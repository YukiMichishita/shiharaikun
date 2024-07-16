package repository

import (
	"context"
	"fmt"
	"shiharaikun/internal/adapter/db/query"
	"shiharaikun/internal/domain/model"
	"shiharaikun/internal/domain/repository"
)

type invoiceRepository struct{}

func NewInvoiceRepository() repository.InvoiceRepository {
	return &invoiceRepository{}
}

func (i *invoiceRepository) CreateInvoice(ctx context.Context, input *model.Invoice) (*model.Invoice, error) {
	invoice := query.Q.Invoice
	if err := invoice.WithContext(ctx).Create(input); err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}
	output, err := invoice.WithContext(ctx).Where(invoice.ID.Eq(input.ID)).First()
	if err != nil {
		return nil, fmt.Errorf("failed to get created invoice: %w", err)
	}
	return output, nil
}

func (i *invoiceRepository) ListInvoicesByDueDate(ctx context.Context, input *repository.ListInvoiceByDueDateInput) ([]*model.Invoice, error) {
	invoice := query.Q.Invoice
	var output []*model.Invoice
	if err := invoice.WithContext(ctx).Where(invoice.DueDate.Between(input.StartDate, input.EndDate), invoice.CompanyID.Eq(input.TenantID)).Scan(output); err != nil {
		return nil, fmt.Errorf("failed to get invoices: %w", err)
	}
	return output, nil
}
