package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func main() {
	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("crm", opts)
	if err != nil {
		panic(err)
	}

	customerService, err := weaver.Get[CustomerService](root)
	if err != nil {
		panic(err)
	}

	invoiceService, err := weaver.Get[InvoiceService](root)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		customers, err := customerService.GetCustomers(r.Context())
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	})

	http.HandleFunc("/invoices", func(w http.ResponseWriter, r *http.Request) {
		invoices, err := invoiceService.GetInvoices(r.Context())
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(invoices)
	})

	http.Serve(lis, nil)
}
