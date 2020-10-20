package routes

import (
	"github.com/gorilla/mux"
	"baddrivers/service"
	"net/http"
	"strconv"
)

// Handler function for all bad driver statuses
func GetBadDriverStatuses(w http.ResponseWriter, r *http.Request) {
	statuses := service.FetchBadDriverStatus()

	if len(statuses) < 1 {
		WriteEmptyArray(w)
		return
	}

	WriteSuccess(w, statuses)
}

// Handler function for getting a bad driver status by it's id
func GetBadDriverStatusById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	status, err := service.FetchBadDriverStatuById(id)

	if err != nil {
		WriteError(w, 500, "internal server error")
		return
	}

	if status.Id == 0 {
		WriteError(w, 404, "not found")
		return
	}

	WriteSuccess(w, status)
}