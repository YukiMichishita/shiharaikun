package model

import "time"

type Invoice struct {
	IssueDate   string `json:"issueDate"`
	AmountDue   int    `json:"amountDue"`
	Fee         int    `json:"fee"`
	FeeRate     string `json:"feeRate"`
	Tax         int    `json:"Tax"`
	TaxRate     string `json:"taxRate"`
	TotalAmount int    `json:"totalAmount"`
	DueDate     string `json:"dueDate"`
	Status      string `json:"status"`
}

type CreateInvoiceRequest struct {
	TenantID  int
	ClientID  int    `json:"clientId"`
	IssueDate string `json:"issueDate"`
	AmountDue int    `json:"amountDue"`
	TaxRate   string `json:"taxRate"`
	DueDate   string `json:"dueDate"`
}

type CreateInvoiceResponse struct {
	Invoice *Invoice `json:"invoice"`
}

type GetInvoicesRequest struct {
	TenantID  int
	StartDate time.Time `uri:"startDate"`
	EndDate   time.Time `uri:"endDate"`
}

type GetInvoicesResponse struct {
	Invoices []*Invoice `json:"invoices"`
}
