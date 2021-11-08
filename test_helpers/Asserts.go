/*
 * Copyright 2019 InsideSales.com Inc.
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

package test_helpers

import (
	"bitbucket.xant.tech/ci/ci-go-utils/common"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

/*
 /$$$$$$$                     /$$
| $$__  $$                   |__/
| $$  \ $$ /$$$$$$   /$$$$$$$ /$$  /$$$$$$$
| $$$$$$$ |____  $$ /$$_____/| $$ /$$_____/
| $$__  $$ /$$$$$$$|  $$$$$$ | $$| $$
| $$  \ $$/$$__  $$ \____  $$| $$| $$
| $$$$$$$/  $$$$$$$ /$$$$$$$/| $$|  $$$$$$$
|_______/ \_______/|_______/ |__/ \_______/
*/
//region BasicTestHelpers

// Asserts that the given expected value matches the given actual value, prints msg if it fails
func AssertEqual(test *testing.T, expected interface{}, actual interface{}, msg string) {
	match := false

	switch actual.(type) {
	case time.Time:
		match = expected.(time.Time).Equal(actual.(time.Time))
	default:
		match = expected == actual
	}

	if !match {
		test.Errorf("%v. Expected '%v', got '%v'", msg, expected, actual)
	}
}

// Asserts that the given expected type matches the given actual type, prints msg if it fails
func AssertSameType(test *testing.T, expected interface{}, actual interface{}, msg string) {
	expected_type := reflect.TypeOf(expected)
	actual_type := reflect.TypeOf(actual)

	if expected_type != actual_type {
		test.Errorf("%v. Expected type: '%v', actual type: '%v'", msg, expected_type, actual_type)
	}
}

//endregion BasicTestHelpers

/*
 /$$$$$$ /$$$$$$$                                /$$   /$$
|_  $$_/| $$__  $$                              | $$  | $$
  | $$  | $$  \ $$  /$$$$$$   /$$$$$$$ /$$   /$$| $$ /$$$$$$
  | $$  | $$$$$$$/ /$$__  $$ /$$_____/| $$  | $$| $$|_  $$_/
  | $$  | $$__  $$| $$$$$$$$|  $$$$$$ | $$  | $$| $$  | $$
  | $$  | $$  \ $$| $$_____/ \____  $$| $$  | $$| $$  | $$ /$$
 /$$$$$$| $$  | $$|  $$$$$$$ /$$$$$$$/|  $$$$$$/| $$  |  $$$$/
|______/|__/  |__/ \_______/|_______/  \______/ |__/   \___/
*/
//region IResultTestHelpers
func BadResult() common.IResult { return common.MakeDefaultCommsResult() }

func GoodResult() common.IResult {
	r := common.MakeDefaultCommsResult()
	r.Succeed()
	return r
}

// Asserts that the given result was successful, prints msg if it fails
func AssertSuccess(test *testing.T, result common.IResult, msg string) {
	if !result.WasSuccessful() {
		if msg == "" {
			msg = "Expected Success, found Failure!"
		}
		test.Error(msg)
	}
	return
}

// Asserts that the given result was successful, prints msg if it fails
func AssertSuccessf(test *testing.T, result common.IResult, msg string, args ...interface{}) {
	if !result.WasSuccessful() {
		if msg == "" {
			msg = "Expected Success, found Failure!"
		}
		test.Errorf(msg, args)
	}
	return
}

// Asserts that the given result was a failure, prints msg if it fails
func AssertFailure(test *testing.T, result common.IResult, msg string) {
	if result.WasSuccessful() {
		if msg == "" {
			msg = "Expected Failure, found Success!"
		}
		test.Error(msg)
	}
	return
}

// Asserts that the given result was a failure, prints msg if it fails
func AssertFailuref(test *testing.T, result common.IResult, msg string, args ...interface{}) {
	if result.WasSuccessful() {
		if msg == "" {
			msg = "Expected Failure, found Success!"
		}
		test.Errorf(msg, args)
	}
	return
}

func AssertResultStatusCodeAndMessage(test *testing.T, result common.IResult, code int, msg string) {
	if c := result.GetStatusCode(); c != code {
		test.Errorf("Incorrect status code found in result! Expected: %d Got: %d", code, c)
	}
	if m := result.GetResponseMessage(); m != msg {
		test.Errorf("Incorrect response message found in result! Expected: %s Got: %s", msg, m)
	}
	return
}

//endregion IResultTestHelpers

/*
 /$$   /$$ /$$$$$$$$/$$$$$$$$/$$$$$$$
| $$  | $$|__  $$__/__  $$__/ $$__  $$
| $$  | $$   | $$     | $$  | $$  \ $$
| $$$$$$$$   | $$     | $$  | $$$$$$$/
| $$__  $$   | $$     | $$  | $$____/
| $$  | $$   | $$     | $$  | $$
| $$  | $$   | $$     | $$  | $$
|__/  |__/   |__/     |__/  |__/

*/
//region HttpTestHelpers

type NopCloser struct {
	io.Reader
}

func (NopCloser) Close() error {
	return nil
}

// Asserts that the given response matches the given http status and response body
func AssertHttpStatusAndMessage(
	test *testing.T,
	response *httptest.ResponseRecorder,
	expected_http_status int,
	expected_response_body string,
) {
	if expected_http_status != response.Code {
		test.Errorf("HTTP code mismatch. Expected '%v', got '%v'", expected_http_status, response.Code)
	}

	if expected_response_body != "" {
		data, err := ioutil.ReadAll(response.Body)
		if err == nil {
			type expected_response_schema struct {
				Error string `json:"error"`
			}
			result := expected_response_schema{}
			json.Unmarshal(data, &result)
			if result.Error != expected_response_body {
				test.Errorf("Error message mismatch. Expected '%v', got '%v'", expected_response_body, result.Error)
			}
		} else {
			test.Error("Failed parsing response body")
		}
	}
}

//Asserts that the http status and 'error'/'message' fields match the given expected values
func AssertHttpStatusAndErrorAndMessage(
	test *testing.T,
	response *httptest.ResponseRecorder,
	expected_http_status int,
	expected_error_response string,
	expected_message_response string,
) {
	if expected_http_status != response.Code {
		test.Errorf("HTTP code mismatch. Expected '%v', got '%v'", expected_http_status, response.Code)
	}

	if expected_error_response != "" || expected_message_response != "" {
		data, err := ioutil.ReadAll(response.Body)
		if err == nil {
			type expected_response_schema struct {
				Error   string `json:"error"`
				Message string `json:"message"`
			}
			result := expected_response_schema{}
			json.Unmarshal(data, &result)
			if result.Error != expected_error_response {
				test.Errorf("Response error mismatch. Expected '%v', got '%v'", expected_error_response, result.Error)
			}
			if result.Message != expected_message_response {
				test.Errorf("Response message mismatch. Expected '%v', got '%v'", expected_message_response, result.Message)
			}
		} else {
			test.Error("Failed parsing response body")
		}
	}
}

//endregion HttpTestHelpers
