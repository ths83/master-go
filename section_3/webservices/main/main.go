package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	app := &App{
		storage: map[string]*Quote{},
	}

	prefix := "/api/v1/"
	http.HandleFunc(prefix+"quote/", app.handleQuote)
	http.HandleFunc(prefix+"quotes/", app.handleQuotesList)

	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Source string `json:"source,omitempty"`
}

type App struct {
	storage map[string]*Quote
}

func (app *App) handleQuote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		app.create(w, r)
	case "GET":
		app.read(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (app *App) create(w http.ResponseWriter, r *http.Request) {
	quote := &Quote{}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		getInternalServerError(w, "An error occurred while reading the value: ", err)
	}

	err = json.Unmarshal(bytes, quote)
	if err != nil {
		getInternalServerError(w, "An error occurred while unmarshalling the value: ", err)
	}

	app.storage[quote.Author] = quote
}

func getInternalServerError(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	io.WriteString(w, message+err.Error())
}

func (app *App) read(w http.ResponseWriter, r *http.Request) {
	key := extractKey(r)
	value, ok := app.storage[key]
	if ok {
		quoteJSON, err := json.Marshal(value)
		if err != nil {
			getInternalServerError(w, "An error occurred while marshalling the value: ", err)
		}

		io.WriteString(w, string(quoteJSON))
	}

	w.WriteHeader(http.StatusNotFound)
}

func extractKey(r *http.Request) string {
	url := strings.Split(r.URL.Path, "/")
	return url[len(url)-1]
}

func (app *App) handleQuotesList(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		app.readAll(w)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (app *App) readAll(w http.ResponseWriter) {
	var quotes []Quote
	for _, value := range app.storage {
		quotes = append(quotes, *value)
	}

	if len(quotes) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	quotesJSON, err := json.Marshal(quotes)
	if err != nil {
		getInternalServerError(w, "An error occurred while marshalling the values: ", err)
	}

	io.WriteString(w, string(quotesJSON))
}
