package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"mentorapi/routes"
	"net/http"
	"os"
)

var Dir string

func SetFlags() {
	flag.StringVar(&Dir, "dir", "./static", "The directory to serve files from, defaults to static")
	flag.Parse()
}

func main() {
	SetFlags()
	_ = godotenv.Load()

	r := mux.NewRouter()
	// Create the home routes
	r.Path("/baddriver").Methods("GET", "POST").HandlerFunc(routes.BadDriverHandler)
	r.Path("/baddriver/{id:[0-9]+}").Methods("GET").HandlerFunc(routes.GetBadDriver)
	r.Path("/baddriverstatus").Methods("GET").HandlerFunc(routes.GetBadDriverStatuses)
	r.Path("/baddriverstatus/{id:[0-9]+}").Methods("GET").HandlerFunc(routes.GetBadDriverStatusById)
	r.Path("/").Methods("GET").HandlerFunc(routes.HomeHandler)

	// Static files
	r.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir(Dir)),
		),
	)

	// Middleware
	r.Use(mux.CORSMethodMiddleware(r)) // deal with CORS

	// Global handlers
	loggedRouter := handlers.CombinedLoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}