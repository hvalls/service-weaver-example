package main

// Code generated by "weaver generate". DO NOT EDIT.
import (
	"context"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"time"
)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "crm/CustomerService",
		Iface: reflect.TypeOf((*CustomerService)(nil)).Elem(),
		New:   func() any { return &customerService{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return customerService_local_stub{impl: impl.(CustomerService), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return customerService_client_stub{stub: stub, getCustomersMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "crm/CustomerService", Method: "GetCustomers"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return customerService_server_stub{impl: impl.(CustomerService), addLoad: addLoad}
		},
	})
	codegen.Register(codegen.Registration{
		Name:  "crm/InvoiceService",
		Iface: reflect.TypeOf((*InvoiceService)(nil)).Elem(),
		New:   func() any { return &invoiceService{} },
		LocalStubFn: func(impl any, tracer trace.Tracer) any {
			return invoiceService_local_stub{impl: impl.(InvoiceService), tracer: tracer}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return invoiceService_client_stub{stub: stub, getInvoicesMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "crm/InvoiceService", Method: "GetInvoices"})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return invoiceService_server_stub{impl: impl.(InvoiceService), addLoad: addLoad}
		},
	})
}

// Local stub implementations.

type customerService_local_stub struct {
	impl   CustomerService
	tracer trace.Tracer
}

func (s customerService_local_stub) GetCustomers(ctx context.Context) (r0 []string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.CustomerService.GetCustomers", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetCustomers(ctx)
}

type invoiceService_local_stub struct {
	impl   InvoiceService
	tracer trace.Tracer
}

func (s invoiceService_local_stub) GetInvoices(ctx context.Context) (r0 []string, err error) {
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "main.InvoiceService.GetInvoices", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.GetInvoices(ctx)
}

// Client stub implementations.

type customerService_client_stub struct {
	stub                codegen.Stub
	getCustomersMetrics *codegen.MethodMetrics
}

func (s customerService_client_stub) GetCustomers(ctx context.Context) (r0 []string, err error) {
	// Update metrics.
	start := time.Now()
	s.getCustomersMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.CustomerService.GetCustomers", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getCustomersMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getCustomersMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.getCustomersMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	if err != nil {
		return
	}
	s.getCustomersMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_string_4af10117(dec)
	err = dec.Error()
	return
}

type invoiceService_client_stub struct {
	stub               codegen.Stub
	getInvoicesMetrics *codegen.MethodMetrics
}

func (s invoiceService_client_stub) GetInvoices(ctx context.Context) (r0 []string, err error) {
	// Update metrics.
	start := time.Now()
	s.getInvoicesMetrics.Count.Add(1)

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "main.InvoiceService.GetInvoices", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
		err = s.stub.WrapError(err)

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			s.getInvoicesMetrics.ErrorCount.Add(1)
		}
		span.End()

		s.getInvoicesMetrics.Latency.Put(float64(time.Since(start).Microseconds()))
	}()

	var shardKey uint64

	// Call the remote method.
	s.getInvoicesMetrics.BytesRequest.Put(0)
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	if err != nil {
		return
	}
	s.getInvoicesMetrics.BytesReply.Put(float64(len(results)))

	// Decode the results.
	dec := codegen.NewDecoder(results)
	r0 = serviceweaver_dec_slice_string_4af10117(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type customerService_server_stub struct {
	impl    CustomerService
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s customerService_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetCustomers":
		return s.getCustomers
	default:
		return nil
	}
}

func (s customerService_server_stub) getCustomers(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetCustomers(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_string_4af10117(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

type invoiceService_server_stub struct {
	impl    InvoiceService
	addLoad func(key uint64, load float64)
}

// GetStubFn implements the stub.Server interface.
func (s invoiceService_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "GetInvoices":
		return s.getInvoices
	default:
		return nil
	}
}

func (s invoiceService_server_stub) getInvoices(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.GetInvoices(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	serviceweaver_enc_slice_string_4af10117(enc, r0)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Encoding/decoding implementations.

func serviceweaver_enc_slice_string_4af10117(enc *codegen.Encoder, arg []string) {
	if arg == nil {
		enc.Len(-1)
		return
	}
	enc.Len(len(arg))
	for i := 0; i < len(arg); i++ {
		enc.String(arg[i])
	}
}

func serviceweaver_dec_slice_string_4af10117(dec *codegen.Decoder) []string {
	n := dec.Len()
	if n == -1 {
		return nil
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = dec.String()
	}
	return res
}