package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

type ErrorMessage struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

var homePage []byte
var homePageSize int

// Load the home page from disk and cache it in memory
// If the file cannot be found, the server crashes.
func getHomePage() ([]byte, int) {
	if homePageSize > 0 {
		return homePage, homePageSize
	}

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal("Fatal Error: Cannot get current working directory.")
	}

	fileName := cwd + "/static/index.html"
	homePage, err = ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal("Could not load index file. We need that to run the API. Please make one.")
	}

	homePageSize = len(homePage)

	return homePage, homePageSize
}

// Send the page to the browser
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	page, size := getHomePage()

	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Content-Length", fmt.Sprint(size))
	w.WriteHeader(200)
	_, err := w.Write(page)

	if err != nil {
		log.Fatal(err)
	}
}

// A utility function to send a JSON API response
// It accepts anything that can be marshalled into JSON
func WriteSuccess(w http.ResponseWriter, v interface {}) {
	body, err := json.Marshal(v)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", fmt.Sprint(len(body)))
	w.WriteHeader(200)
	_, err = w.Write(body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error sending bad drivers: %v\n", err)
	}
}

// A utility function to write an error message for the API
// Accepts the status code and an error message. It should go without saying, but avoid sending actual debugging/error
// data. Instead create a generic error message that a developer will identify, but gives specifics about implementation.
func WriteError(w http.ResponseWriter, status int, message string) {
	var error ErrorMessage
	error.Status = status
	error.Message = message

	body, err := json.Marshal(error)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("internal server error"))
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Length", fmt.Sprint(len(body)))
	w.WriteHeader(status)
	w.Write(body)
}

// Write an empty array response
// Use this when a database response returns no results so that you can still send a valid 200 response with an array,
// but don't have any results.
func WriteEmptyArray(w http.ResponseWriter) {
	WriteSuccess(w, make([]string, 0))
}