package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := getRouter()

	log.Println("Listening...")
	err := http.ListenAndServe(":7777", r)
	if err != nil {
		log.Fatal(err)
	}
}

func getRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/echo/{message}", Echo).Methods("GET")
	return r
}

// Echo responds to the /api/v1/echo/ endpoint
func Echo(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s for %s\n", r.Method, r.URL.Path)

	say := mux.Vars(r)["message"]
	w.Write([]byte(say))
}
