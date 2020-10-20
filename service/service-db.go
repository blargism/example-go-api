package service

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"os"
)

var conn *pgx.Conn

func Conn() *pgx.Conn {
	fmt.Print(conn)
	if conn != nil {
		return conn
	}
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	host := os.Getenv("PGHOST")
	database := os.Getenv("PGDATABASE")

	connectionString := fmt.Sprintf("postgres://%v:%v@%v/%v", user, password, host, database)

	var err error
	conn, err = pgx.Connect(context.Background(), connectionString)

	if os.IsExist(err) {
		log.Fatalf("Cannot connect to %v/%v", host, database)
	}

	return conn
}

func RunQuery(query string, params ...interface {}) (pgx.Rows, error) {
	rows, err := Conn().Query(context.Background(), query, params...)
	return rows, err
}

func RunSingle(query string, params ...interface {}) (pgx.Row) {
	rows := Conn().QueryRow(context.Background(), query, params...)

	return rows
}

func RunSimpleQuery(query string) (pgx.Rows, error) {
	rows, err := Conn().Query(context.Background(), query)
	return rows, err
}