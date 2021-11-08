package routing

import (
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/propagation"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/codes"
	"net/http"
	"time"
	"context"
)

const CONTEXT_API_NAME = "api_name"

type JaegerHTTPClientWrapper struct {
	r http.RoundTripper
}

func (t JaegerHTTPClientWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	operationName, _ := req.Context().Value(CONTEXT_API_NAME).(string)
	if operationName == "" {
		operationName = req.URL.Host
	}

	tracer := global.Tracer("")
	ctx := req.Context()
	ctx, span := tracer.Start(ctx, operationName, trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	span.SetAttribute(string("http.url"), req.URL.String())
	span.SetAttribute(string("http.method"), req.Method)

	propagation.InjectHTTP(ctx, global.Propagators(), req.Header)
	req = req.WithContext(trace.ContextWithSpan(ctx, span))

	base := t.r
	if base == nil {
		base = http.DefaultTransport
	}
	resp, err := base.RoundTrip(req)
	if err == nil {
		span.SetAttribute("http.status_code", resp.StatusCode)
		if resp.StatusCode >= http.StatusBadRequest {
			span.SetAttribute("error", true)
		}
		for header, value := range resp.Header {
			span.SetAttribute("http.response.header."+header, value)
		}
		span.SetStatus(getTraceCode(resp.StatusCode), "")
	} else {
		span.SetAttribute("error.desc", err.Error())
		span.SetAttribute("error", true)
		span.SetStatus(codes.Unknown, err.Error())
	}

	return resp, err
}

// CloneRequest creates a shallow copy of the request along with a deep copy of the Headers.
func CloneRequest(req *http.Request) *http.Request {
	r := new(http.Request)

	// shallow clone
	*r = *req

	// deep copy headers
	r.Header = CloneHeader(req.Header)

	return r
}

// CloneHeader creates a deep copy of an http.Header.
func CloneHeader(in http.Header) http.Header {
	out := make(http.Header, len(in))
	for key, values := range in {
		newValues := make([]string, len(values))
		copy(newValues, values)
		out[key] = newValues
	}
	return out
}

//This struct allows tracing info to be passed into async api calls without the worry of the context being canceled
type ValueOnlyContext struct{ context.Context }

func (ValueOnlyContext) Deadline() (deadline time.Time, ok bool) { return }
func (ValueOnlyContext) Done() <-chan struct{}                   { return nil }
func (ValueOnlyContext) Err() error                              { return nil }
