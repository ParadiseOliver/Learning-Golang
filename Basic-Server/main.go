package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellowHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func competitionHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %w", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	competition := r.FormValue("competition")
	place := r.FormValue("position")
	fmt.Fprintf(w, "Competition: %s\n", competition)
	fmt.Fprintf(w, "Position: %s", place)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/competition", competitionHandler)
	http.HandleFunc("/hello", hellowHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
