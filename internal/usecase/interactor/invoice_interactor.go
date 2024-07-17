package interactor

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"shiharaikun/internal/domain/entity"
	domainmodel "shiharaikun/internal/domain/model"
	"shiharaikun/internal/domain/repository"
	"shiharaikun/internal/usecase"
	"shiharaikun/internal/usecase/model"
	"time"
)

type invoiceInterActor struct {
	repo repository.InvoiceRepository
}

func NewInvoiceInterActor(repo repository.InvoiceRepository) usecase.InvoiceUseCase {
	return &invoiceInterActor{repo: repo}
}

func (i *invoiceInterActor) CreateInvoice(ctx context.Context, input *model.CreateInvoiceRequest) (*model.CreateInvoiceResponse, error) {
	amountDue := entity.InvoiceAmount(decimal.NewFromInt(int64(input.AmountDue)))
	fee, err := amountDue.CalcFee()
	if err != nil {
		return nil, fmt.Errorf("failed to calculate fee: %w", err)
	}
	tax, err := amountDue.CalcTax(input.TaxRate)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate tax: %w", err)
	}
	amount, err := amountDue.CalcTotalAmount(input.TaxRate)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate total amount: %w", err)
	}
	tr, err := decimal.NewFromString(input.TaxRate)
	if err != nil {
		return nil, fmt.Errorf("failed to convert tax rate string to decimal: %w", err)
	}
	fr, err := entity.FeeRate()
	if err != nil {
		return nil, fmt.Errorf("failed to get fee rate: %w", err)
	}

	parsedIssuedDate, err := time.Parse(time.DateOnly, input.IssueDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse issued date: %w", err)
	}
	parsedDueDate, err := time.Parse(time.DateOnly, input.DueDate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse due date: %w", err)
	}
	user, ok := ctx.Value("user").(*domainmodel.User)
	if !ok {
		return nil, fmt.Errorf("failed to get user from context")
	}
	invoice := &domainmodel.Invoice{
		CompanyID:      user.CompanyID,
		ClientID:       int32(input.ClientID),
		IssuedDate:     parsedIssuedDate,
		AmountDue:      decimal.Decimal(amountDue),
		Fee:            *fee,
		FeeRate:        *fr,
		ConsumptionTax: *tax,
		TaxRate:        tr,
		TotalAmount:    *amount,
		DueDate:        parsedDueDate,
		// TODO: enumにする
		Status: "pending",
	}

	created, err := i.repo.CreateInvoice(ctx, invoice)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}
	resp := &model.CreateInvoiceResponse{
		Invoice: &model.Invoice{
			IssueDate:   created.IssuedDate.String(),
			AmountDue:   int(created.AmountDue.Round(0).IntPart()),
			Fee:         int(created.Fee.Round(0).IntPart()),
			FeeRate:     created.FeeRate.Mul(decimal.NewFromInt(100)).String(),
			Tax:         int(created.ConsumptionTax.Round(0).IntPart()),
			TaxRate:     created.TaxRate.Mul(decimal.NewFromInt(100)).String(),
			TotalAmount: int(created.TotalAmount.Round(0).IntPart()),
			DueDate:     created.DueDate.String(),
			Status:      created.Status,
		},
	}
	return resp, nil
}

func (i *invoiceInterActor) GetInvoices(ctx context.Context, input *model.GetInvoicesRequest) (*model.GetInvoicesResponse, error) {
	// TODO:implement
	return nil, nil
}
