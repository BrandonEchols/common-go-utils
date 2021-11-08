package routing_test

import (
	"github.com/BrandonEchols/common-go-utils/routing"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthControllerReturns200WithAlive(test *testing.T) {
	mockCtrl := gomock.NewController(test)
	defer mockCtrl.Finish()

	health_controller := routing.GetHealthController()
	req, err := http.NewRequest("GET", "http://localhost:8999/health", nil)
	response := httptest.NewRecorder()
	if err == nil {
		health_controller.GetHealth(response, req)
		if response.Code != 200 {
			test.Errorf("HTTP code mismatch. Expected '%v', got '%v'", 200, response.Code)
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			test.Error("Test did not complete. Message: " + err.Error())
		}

		data_string := string(data[:])
		if data_string != "Alive" {
			test.Error("Expected message 'Alive' got " + data_string)
		}
	} else {
		test.Error("Test did not complete. Message: " + err.Error())
	}
}
