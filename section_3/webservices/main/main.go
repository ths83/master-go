package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	prefix := "/api/v1/"
	http.HandleFunc(prefix+"quote/", handleQuote)
	http.HandleFunc(prefix+"quotes/", handleQuotesList)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}

func handleQuote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		io.WriteString(w, "Create: "+r.URL.Path+"n")
	case "GET":
		io.WriteString(w, "Read: "+r.URL.Path+"n")
	case "PUT":
		io.WriteString(w, "Update: "+r.URL.Path+"n")
	case "DELETE":
		io.WriteString(w, "Delete: "+r.URL.Path+"n")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func handleQuotesList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		io.WriteString(w, "Read: "+r.URL.Path+"n")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
