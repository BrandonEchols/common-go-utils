package test_helpers

import (
	"encoding/json"
	"github.com/BrandonEchols/common-go-utils/common"
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
func BadResult() common.IResult { return common.MakeDefaultCommonResult() }

func GoodResult() common.IResult {
	r := common.MakeDefaultCommonResult()
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
