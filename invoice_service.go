package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type InvoiceService interface {
	GetInvoices(context.Context) ([]string, error)
}

type invoiceService struct {
	weaver.Implements[InvoiceService]
}

func (r *invoiceService) GetInvoices(_ context.Context) ([]string, error) {
	return []string{"Invoice 1", "Invoice 2", "Invoice 3"}, nil
}
