package main

import (
	"io"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
	io.WriteString(w, "URL: "+r.URL.Path)
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
