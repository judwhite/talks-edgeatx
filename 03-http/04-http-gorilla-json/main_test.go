package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestServer() *httptest.Server {
	// set up httptest test server with our router
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router := getRouter()
		router.ServeHTTP(w, r)
	}))
}

func get(ts *httptest.Server, path string) (string, int, error) {
	// call the test server endpoint
	resp, err := http.Get(ts.URL + path)
	if err != nil {
		return "", 0, err
	}

	// read the response
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", 0, err
	}

	return string(body), resp.StatusCode, nil
}

func TestEcho(t *testing.T) {
	// Arrange
	ts := newTestServer()
	defer ts.Close()

	// set up test table (path, expected body, expected HTTP status)
	testData := []struct {
		path           string
		expectedBody   string
		expectedStatus int
	}{
		{"/api/v1/echo/hello!", "{\"message\":\"hello!\"}\n", 200},
		{"/api/v1/echo/hello", "I don't want to say that.", 400},
	}

	for _, test := range testData {
		// Act
		body, status, err := get(ts, test.path)
		if err != nil {
			t.Error(err)
		}

		// Assert
		if body != test.expectedBody {
			t.Errorf("expected: %q actual: %q", test.expectedBody, body)
		}
		if status != test.expectedStatus {
			t.Errorf("expected: %d actual: %d", test.expectedStatus, status)
		}
	}
}
