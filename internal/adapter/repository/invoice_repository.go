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
	iq := query.Q.Invoice
	if err := iq.WithContext(ctx).Create(input); err != nil {
		return nil, fmt.Errorf("failed to create iq: %w", err)
	}
	output, err := iq.WithContext(ctx).Where(iq.ID.Eq(input.ID)).First()
	if err != nil {
		return nil, fmt.Errorf("failed to get created iq: %w", err)
	}
	return output, nil
}

func (i *invoiceRepository) ListInvoicesByDueDate(ctx context.Context, input *repository.ListInvoiceByDueDateInput) ([]*model.Invoice, error) {
	iq := query.Q.Invoice
	var output []*model.Invoice
	if err := iq.WithContext(ctx).Where(iq.DueDate.Between(input.StartDate, input.EndDate), iq.CompanyID.Eq(input.CompanyID)).Scan(output); err != nil {
		return nil, fmt.Errorf("failed to get invoices: %w", err)
	}
	return output, nil
}
