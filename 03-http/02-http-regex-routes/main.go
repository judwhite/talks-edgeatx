package main

import (
	"log"
	"net/http"
	"regexp"
)

// regexEcho specifies the pattern used to handle requests to /api/v1/echo/{msg}
var regexEcho = regexp.MustCompile("^/api/v1/echo/(.+)$")

func main() {
	http.HandleFunc("/api/v1/echo/", Echo)

	log.Println("Listening...")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Echo responds to the /api/v1/echo endpoint
func Echo(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s for %s\n", r.Method, r.URL.Path)

	// only accept GET requests
	if r.Method != http.MethodGet {
		writeStatus(w, 404)
		return
	}

	// extract matches from the regex
	// example: []string{"/api/v1/echo/hi", "hi"}
	matches := regexEcho.FindStringSubmatch(r.URL.Path)

	if len(matches) != 2 {
		// if nothing after /api/v1/echo/ was passed
		writeStatus(w, 404)
		return
	}

	msg := matches[1]

	// echo the message
	w.Write([]byte(msg))
}

func writeStatus(w http.ResponseWriter, statusCode int) {
	// write http status header and text to the body
	w.WriteHeader(statusCode)
	msg := http.StatusText(statusCode)
	w.Write([]byte(msg))
}
