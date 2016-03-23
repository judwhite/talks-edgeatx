package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// handlerFunc takes an http.Request and returns an object to be serialized,
// an HTTP status code, and an error.
//
// If the returned object is a raw string or []byte it will be written
// without serialization.
type handlerFunc func(*http.Request) (interface{}, int, error)

// echoResponse is the response to the /api/v1/echo endpoint.
type echoResponse struct {
	Message string `json:"message"`
}

func main() {
	// set up routing
	r := getRouter()

	// start HTTP server with our router
	log.Println("Listening...")
	err := http.ListenAndServe(":7777", r)
	if err != nil {
		log.Fatal(err)
	}
}

func getRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/echo/{message}", Log(JSON(Echo))).Methods("GET")
	return r
}

// Echo responds to the /api/v1/echo endpoint
func Echo(r *http.Request) (interface{}, int, error) {
	say := mux.Vars(r)["message"]
	if say == "hello" {
		return "I don't want to say that.", 400, nil
	}
	return echoResponse{Message: say}, 200, nil
}

// Log logs the method and URL path
func Log(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s for %s\n", r.Method, r.URL.Path)
		handler(w, r)
	}
}

// JSON wraps a handlerFunc function and marshals the response to JSON.
func JSON(handler handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, status, err := handler(r)

		// log errors
		if err != nil {
			log.Println(err)
		}

		// write non-200 response headers
		if status <= 0 {
			status = 500
		}
		if status != 200 {
			w.WriteHeader(status)
		}

		if resp != nil {
			// if the response is a string or []byte write it directly.
			// otherwise, return a JSON object.
			if str, ok := resp.(string); ok {
				w.Write([]byte(str))
			} else if bytes, ok := resp.([]byte); ok {
				w.Write(bytes)
			} else {
				enc := json.NewEncoder(w)
				enc.Encode(resp)
			}
		}
	}
}
