package usecase

import (
	"context"
	"shiharaikun/internal/usecase/model"
)

type InvoiceUseCase interface {
	CreateInvoice(ctx context.Context, input *model.CreateInvoiceRequest) (*model.CreateInvoiceResponse, error)
	GetInvoices(ctx context.Context, input *model.GetInvoicesRequest) (*model.GetInvoicesResponse, error)
}
