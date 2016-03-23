package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", helloWorld)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s for %s\n", r.Method, r.URL.Path)
	w.Write([]byte("Hello, world!"))
}
