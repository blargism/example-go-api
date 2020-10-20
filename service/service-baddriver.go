package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"strconv"
)

type BadDriver struct {
	Id              int             `json:"id" schema:"id"`
	Name            string          `json:"name" schema:"name"`
	Reason          string          `json:"reason" schema:"reason"`
	Status          BadDriverStatus `json:"status" schema:"status"`
	AccidentCount   int             `json:"accident_count" schema:"accident_count"`
	TicketCount     int             `json:"ticket_count" schema:"ticket_count"`
	KarensIrritated int             `json:"karens_irritated" schema:"karens_irritated"`
}

// Fetch all bad drivers, no limits
func FetchBadDrivers() ([]BadDriver, error) {
	var badDrivers []BadDriver
	rows, err := Conn().Query(context.Background(), "SELECT * FROM bad_drivers bd INNER JOIN bad_driver_statuses bds ON bd.status = bds.bad_driver_status_id")
	badDrivers = extractRecords(rows)

	if err != nil {
		return badDrivers, err
	}

	return badDrivers, err
}

// Fetch bad drivers with an SQL offset (start) and limit.
func FetchBadDriversWithLimits(start int, limit int) ([]BadDriver, error) {
	var badDrivers []BadDriver

	baseQuery := fmt.Sprintf("SELECT * FROM bad_drivers bd INNER JOIN bad_driver_statuses bds ON bd.status = bds.bad_driver_status_id OFFSET %v", start)

	if limit > 0 {
		baseQuery = fmt.Sprintf("%v LIMIT %v", baseQuery, limit)
	}

	rows, err := Conn().Query(context.Background(), baseQuery)
	badDrivers = extractRecords(rows)

	if err != nil {
		return badDrivers, err
	}

	return badDrivers, err
}

// Fetch a single bad driver by bad_driver_id
func FetchBadDriver(idstr string) (badDriver BadDriver, err error) {
	var badDriverStatus BadDriverStatus
	var status int
	id, err := strconv.Atoi(idstr)

	row := Conn().QueryRow(
		context.Background(),
		"SELECT * FROM bad_drivers bd INNER JOIN bad_driver_statuses bds ON bd.status = bds.bad_driver_status_id WHERE bd.bad_driver_id = $1",
		id)

	err = row.Scan(
		&badDriver.Id,
		&badDriver.Name,
		&badDriver.Reason,
		&status,
		&badDriver.AccidentCount,
		&badDriver.TicketCount,
		&badDriver.KarensIrritated,
		&badDriverStatus.Id,
		&badDriverStatus.Title,
		&badDriverStatus.Description)

	if err != nil {
		return badDriver, errors.New("database error")
	}

	badDriver.Status = badDriverStatus

	return badDriver, err
}

// Create a new bad driver in the database
func InsertBadDriver(name *string, reason *string, status *int, accident_count *int, ticket_count *int, karens_irritated *int) (BadDriver, error) {
	var badDriver BadDriver
	var id int

	row:= Conn().QueryRow(
		context.Background(),
		"INSERT INTO bad_drivers (name, reason, status, accident_count, ticket_count, karens_irritated) VALUES ($1, $2, $3, $4, $5, $6) RETURNING bad_driver_id",
		name, reason, status, accident_count, ticket_count, karens_irritated)

	err := row.Scan(&id)

	if err != nil {
		return badDriver, err
	}

	badDriverStatus, err := FetchBadDriverStatuById(*status)

	badDriver.Id = id
	badDriver.Name = *name
	badDriver.Reason = *reason
	badDriver.Status = badDriverStatus
	badDriver.AccidentCount = *accident_count
	badDriver.TicketCount = *ticket_count
	badDriver.KarensIrritated = *karens_irritated

	return badDriver, err
}

// A utility function that extracts bad drivers from database results
func extractRecords(rows pgx.Rows) []BadDriver {
	var badDrivers []BadDriver

	for rows.Next() {
		var badDriver BadDriver
		var badDriverStatus BadDriverStatus
		var status int

		rows.Scan(&badDriver.Id, &badDriver.Name, &badDriver.Reason, &status, &badDriver.AccidentCount, &badDriver.TicketCount, &badDriver.KarensIrritated, &badDriverStatus.Id, &badDriverStatus.Title, &badDriverStatus.Description)

		badDriver.Status = badDriverStatus
		badDrivers = append(badDrivers, badDriver)
	}

	return badDrivers
}
