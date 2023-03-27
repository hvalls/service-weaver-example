package main

import (
	"context"

	"github.com/ServiceWeaver/weaver"
)

type CustomerService interface {
	GetCustomers(context.Context) ([]string, error)
}

type customerService struct {
	weaver.Implements[CustomerService]
}

func (r *customerService) GetCustomers(_ context.Context) ([]string, error) {
	return []string{"Customer 1", "Customer 2", "Customer 3"}, nil
}
