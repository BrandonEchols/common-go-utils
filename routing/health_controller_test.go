/**
 * Copyright 2017 InsideSales.com Inc.
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is the property of InsideSales.com, Inc. and its suppliers, if
 * any. The intellectual and technical concepts contained herein are proprietary and are protected by
 * trade secret or copyright law, and may be covered by U.S. and foreign patents and patents pending.
 * Dissemination of this information or reproduction of this material is strictly forbidden without prior
 * written permission from InsideSales.com Inc.
 *
 * Requests for permission should be addressed to the Legal Department, InsideSales.com,
 * 1712 South East Bay Blvd. Provo, UT 84606.
 *
 * The software and any accompanying documentation are provided "as is" with no warranty.
 * InsideSales.com, Inc. shall not be liable for direct, indirect, special, incidental, consequential, or other
 * damages, under any theory of liability.
 */

package routing_test

import (
	"bitbucket.xant.tech/ci/ci-go-utils/routing"
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
