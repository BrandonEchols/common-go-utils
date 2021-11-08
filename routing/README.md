# A Collection Of Common Go Utilities

## Purpose of this module
This module should contain the classes that are used for web-based-services that will be hosting API's of some kind. 
The following are descriptions of the current files/classes that are available in this module.

#### CustomRouter.go
This contains a ready to use extension of the Goriila/mux router (http://www.gorillatoolkit.org/pkg/mux) that includes
the capability of mounting middleware functions to a group of routes at a time. This router is fully compatible with all
of the middleware defined in the module.

#### JwtAuthenticator.go
This contains a JSON Web Token Authenticator/creator. See the JwtAuthenticator.go for more information.

#### HealthController.go
This is a simple implementation of a 'controller' class that has one http.HandlerFunc GetHealth. This class is used to
host this simple endpoint for purposes of determining if a service is up and routing is working.

#### PrometheusMiddleware.go
This is a wrapper to the prometheus go client (https://github.com/prometheus/client_golang). It wraps the functionality
of prometheus in a middleware that is compatible with the CustomRouter.

#### LoggingMiddleware.go
This is a simplistic middleware that logs out information about the incoming request before sending it through, and then
logs the status code upon return.

#### CORSMiddleware.go
This is a piece of middleware that Handles CORS restriction setting up. See the file for more information.
