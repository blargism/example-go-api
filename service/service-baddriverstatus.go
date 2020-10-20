package service

import (
	"fmt"
	"os"
)

type BadDriverStatus struct {
	Id int `json:"id" schema:"id"`
	Title string `json:"title" schema:"title"`
	Description string `json:"description" schema:"description"`
}

// Fetches all bad driver statuses from the database
func FetchBadDriverStatus() []BadDriverStatus {
	var statuses []BadDriverStatus
	rows, err := RunSimpleQuery("SELECT * FROM bad_driver_statuses")

	if os.IsExist(err) {
		fmt.Errorf("error running the query: %v", err)
	}

	for rows.Next() {
		var bds BadDriverStatus
		rows.Scan(&bds.Id, &bds.Title, &bds.Description)
		statuses = append(statuses, bds)
	}

	return statuses
}

// Fetches a single bad driver status by it's id
func FetchBadDriverStatuById(id int) (BadDriverStatus, error) {
	var status BadDriverStatus
	row := RunSingle("SELECT * FROM bad_driver_statuses WHERE bad_driver_status_id = $1", id)

	err := row.Scan(&status.Id, &status.Title, &status.Description)

	return status, err
}