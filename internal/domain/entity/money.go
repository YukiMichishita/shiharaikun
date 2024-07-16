package entity

import (
	"fmt"
	"github.com/shopspring/decimal"
)

const FeeRateString = "0.4"

type Money decimal.Decimal

type InvoiceAmount Money

func (i InvoiceAmount) CalcFee() (*decimal.Decimal, error) {
	fr, err := decimal.NewFromString(FeeRateString)
	if err != nil {
		return nil, fmt.Errorf("failed to convert fee rate string to decimal: %w", err)
	}
	fee := fr.Mul(decimal.Decimal(i))
	return &fee, nil
}

func (i InvoiceAmount) CalcTax(rate string) (*decimal.Decimal, error) {
	taxRate, err := decimal.NewFromString(rate)
	if err != nil {
		return nil, fmt.Errorf("failed to convert tax rate string to decimal: %w", err)
	}
	tax := taxRate.Mul(decimal.Decimal(i))
	return &tax, nil
}

func (i InvoiceAmount) CalcTotalAmount(taxRatePercent string) (*decimal.Decimal, error) {
	fee, err := i.CalcFee()
	if err != nil {
		return nil, fmt.Errorf("failed to calculate fee: %w", err)
	}
	tax, err := i.CalcTax(taxRatePercent)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate tax: %w", err)
	}
	totalAmount := decimal.Decimal(i).Add(*fee).Add(*tax)
	return &totalAmount, nil
}

func FeeRate() (*decimal.Decimal, error) {
	fr, err := decimal.NewFromString(FeeRateString)
	if err != nil {
		return nil, fmt.Errorf("failed to convert fee rate string to decimal: %w", err)
	}
	return &fr, nil
}
