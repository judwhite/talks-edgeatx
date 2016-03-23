package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEcho(t *testing.T) {
	// Arrange

	// set up httptest test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router := getRouter()
		router.ServeHTTP(w, r)
	}))
	defer ts.Close()

	// Act

	// call the test server endpoint /api/v1/echo/hello!
	resp, err := http.Get(ts.URL + "/api/v1/echo/hello!")
	if err != nil {
		t.Fatal(err)
	}

	// read the response
	actual, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	// Assert

	expected := []byte("hello!")
	if bytes.Compare(actual, expected) != 0 {
		t.Fatalf("expected: %s actual: %s", expected, actual)
	}
}
