package routing

import (
	"net/http"
)

/*
	This class is used for hosting a 'health' point that is useful for determining if the service is up and routing
	is working as expected.
*/
type IHealthController interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
}

//Implements IHealthController
type healthController struct{}

/*
	Returns an implementation of IHealthController
*/
func GetHealthController() IHealthController {
	return &healthController{}
}

/*
	This API endpoint simply returns a UTF-8 character response: "Alive".
*/
func (this *healthController) GetHealth(w http.ResponseWriter, r *http.Request) {
	result := []byte("Alive")

	w.Header().Set("Content-Type", "charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
