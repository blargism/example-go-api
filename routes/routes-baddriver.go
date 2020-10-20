package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"io/ioutil"
	"baddrivers/service"
	"net/http"
	"os"
	"strconv"
)

type BadDriverPost struct {
	Name string `json:"name" schema:"name"`
	Reason string `json:"reason" schema:"reason"`
	Status int `json:"status" schema:"status"`
	AccidentCount int `json:"accident_count" schema:"accident_count"`
	TicketCount int `json:"ticket_count" schema:"ticket_count"`
	KarensIrritated int `json:"karens_irritated" schema:"karens_irritated"`
}

var decoder = schema.NewDecoder()

// Route to the proper bad driver handler based on request method
func BadDriverHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		CreateBadDriver(w, r)
	} else {
		GetBadDrivers(w, r)
	}
}

// Get a list of bad drivers
// Can be paginated via start and limit query parameters
func GetBadDrivers(w http.ResponseWriter, r *http.Request) {
	var badDrivers []service.BadDriver

	query := r.URL.Query()
	var start int
	var limit int
	var constrain bool
	constrain = false

	if len(query["start"]) > 0 {
		start, _ = strconv.Atoi(query["start"][0])
		constrain = true
	}

	if len(query["limit"]) > 0 {
		limit, _ = strconv.Atoi(query["limit"][0])
		constrain = true
	}

	if constrain {
		badDrivers, _ = service.FetchBadDriversWithLimits(start, limit)
	} else {
		badDrivers, _ = service.FetchBadDrivers()
	}

	if len(badDrivers) < 1 {
		WriteEmptyArray(w)
		return
	}

	WriteSuccess(w, badDrivers)
}

// Get a bad driver by it's id
func GetBadDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	badDriver, err := service.FetchBadDriver(id)

	if err != nil {

		WriteError(w, 500, "server error")
		return;
	}

	if badDriver.Id == 0 {
		WriteError(w,404, "record not found")
		return
	}

	WriteSuccess(w, badDriver)
}

// Create a new bad driver in the database
func CreateBadDriver(w http.ResponseWriter, r *http.Request) {
	var bd BadDriverPost
	var err error
	if r.Header.Get("Content-Type") == "application/json" {
		body, _ := ioutil.ReadAll(r.Body)
		err = json.Unmarshal(body, &bd)
	} else {
		WriteError(w, 400, "requires json content type")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "error parisng the response")
		WriteError(w, 400,"bad request")
		return
	}

	badDriver, err := service.InsertBadDriver(
		&bd.Name,
		&bd.Reason,
		&bd.Status,
		&bd.AccidentCount,
		&bd.TicketCount,
		&bd.KarensIrritated)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error parisng the response")
		WriteError(w, 400, "bad request")
		return
	}

	WriteSuccess(w, badDriver)
}
