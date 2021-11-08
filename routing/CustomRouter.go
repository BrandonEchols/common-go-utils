package routing

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/propagation"
	"go.opentelemetry.io/otel/api/trace"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

//The Route struct is used by the service implementing this router to use with the RegisterRoute function
type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

//The Path struct is used by the service implementing this router to use with the RegisterPath function
type Path struct {
	Path        string
	Prefix      string
	HandlerFunc http.HandlerFunc
}

//The CustomRouter struct extends Gorilla's Mux Router to allow for customizing a router
//Additional functions/data members can be added to CustomRouter if you want additional functionality
type CustomRouter struct {
	*mux.Router
	mw         []MiddleWare
	ParentPath string
}

func GetCustomRouter(r *mux.Router, m []MiddleWare) *CustomRouter {
	return &CustomRouter{
		Router: r,
		mw:     m,
	}
}

func (r *CustomRouter) GetCustomSubRouter(path string) *CustomRouter {
	sub_router := CustomRouter{Router: r.PathPrefix(path).Subrouter()}
	sub_router.ParentPath = r.ParentPath + path
	return &sub_router
}

//GetSubRouterWithMiddleware is used to get a subrouter (on the given path) that uses the same middleware as the parent
//CustomRouter. If you then want to add additional middleware to the subrouter, see the 'Use' function
func (r *CustomRouter) GetSubRouterWithMiddleware(path string) *CustomRouter {
	sub_router := CustomRouter{Router: r.PathPrefix(path).Subrouter()}
	sub_router.Use(r.mw)
	return &sub_router
}

//MiddleWare is a function interface that is implemented by any function that wraps an http.HandlerFunc
type MiddleWare func(http.HandlerFunc) http.HandlerFunc

//The Use function allows middleware to be added to the router dynamically. You can add any number MiddleWare functions
//This middleware will be applied to any endpoint/route registered using the RegisterRoute function
func (r *CustomRouter) Use(middle_ware []MiddleWare) {
	for _, mw := range middle_ware {
		r.mw = append(r.mw, mw)
	}
}

//The ClearMiddleware function removes all middleware from the CustomRouter
func (r *CustomRouter) ClearMiddleware() {
	r.mw = []MiddleWare{}
}

//The RegisterRoute function is a custom function that wraps the passed in route in any middleware that has been
//added to the router
//The order that the middleware will be applied to an endpoint is Last In First Out
//For example: The first MiddleWare in the array passed into the Use function, will be the last middleware
//applied to a given route
func (r *CustomRouter) RegisterRoute(route Route) {
	handler := route.HandlerFunc
	for _, mw := range r.mw {
		handler = mw(handler)
	}

	r.Methods(route.Method).Path(route.Path).Handler(handler)
}

//The RegisterPath function is a custom function that wraps the passed in path in any middleware that has been
//added to the router and will include anything underneath that path (see mux.router.PathPrefix for more information)
//The order that the middleware will be applied to an endpoint is Last In First Out
//For example: The first MiddleWare in the array passed into the Use function, will be the last middleware
//applied to a given route
func (r *CustomRouter) RegisterPath(path Path) {
	handler := path.HandlerFunc
	for _, mw := range r.mw {
		handler = mw(handler)
	}

	r.PathPrefix(path.Path).Handler(
		http.StripPrefix(path.Prefix,
			handler,
		),
	)
}

//The RegisterRouteWithTracing function is a custom function that wraps the passed in route in any middleware that has been
//added to the router with the addition of tracing middleware at the outer most level
//The order that the middleware will be applied to an endpoint is Last In First Out
func (r *CustomRouter) RegisterRouteWithTracing(route Route) {
	handler := route.HandlerFunc
	for _, mw := range r.mw {
		handler = mw(handler)
	}
	//Add Tracing Middleware
	handler = r.OpenTelemetryMiddleware(route, handler)

	r.Methods(route.Method).Path(route.Path).Handler(handler)
}

func (router *CustomRouter) OpenTelemetryMiddleware(route Route, h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Trace the request
		tracer := global.Tracer("")
		ctx := propagation.ExtractHTTP(r.Context(), global.Propagators(), r.Header)
		ctx, span := tracer.Start(ctx, fmt.Sprintf("%s %s", r.Method, fmt.Sprintf("%s %s%s", r.Method, router.ParentPath, route.Path)),
			trace.WithSpanKind(trace.SpanKindServer))
		defer span.End()

		span.SetAttribute(string("http.url"), r.URL.String())
		span.SetAttribute(string("http.method"), r.Method)

		propagation.InjectHTTP(ctx, global.Propagators(), r.Header)
		r = r.WithContext(trace.ContextWithSpan(ctx, span))

		//Wrap the response writer to get the status code
		wrapper := &HTTPResponseWriterWrapper{ResponseWriter: w}

		h.ServeHTTP(wrapper, r.WithContext(ctx))

		if wrapper.status_code != 0 {
			span.SetStatus(getTraceCode(wrapper.status_code), "")
		}
	})
}

type HTTPResponseWriterWrapper struct {
	http.ResponseWriter
	status_code int
}

func (this *HTTPResponseWriterWrapper) WriteHeader(code int) {
	this.status_code = code
	this.ResponseWriter.WriteHeader(code)
}

func getTraceCode(respCode int) codes.Code {
	switch respCode {
	case 400:
		return codes.InvalidArgument
	case 504:
		return codes.DeadlineExceeded
	case 404:
		return codes.NotFound
	case 403:
		return codes.PermissionDenied
	case 401:
		return codes.Unauthenticated
	case 429:
		return codes.ResourceExhausted
	case 501:
		return codes.Unimplemented
	case 503:
		return codes.Unavailable
	default:
		if respCode < 400 && respCode >= 200 {
			return codes.OK
		}
		return codes.Unknown
	}
}
