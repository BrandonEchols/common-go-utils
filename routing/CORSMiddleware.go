package routing

import (
	"github.com/gorilla/handlers"
	"net/http"
)

/*
	CorsMiddleware is used for Allowing Cross-origin requests.
	@params
		h http.Handler The handler to wrap
		origins []string An array of origins to allow
		methods []string An array of methods to allow
		headers []string An array of headers to allow
*/
func CorsMiddleware(h http.Handler, origins []string, methods []string, headers []string) http.Handler {
	originsOk := handlers.AllowedOrigins(origins)
	methodsOk := handlers.AllowedMethods(methods)
	headersOk := handlers.AllowedHeaders(headers)

	return handlers.CORS(originsOk, methodsOk, headersOk)(h)
}
