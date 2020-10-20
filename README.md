# example-go-api

An example Go API that connects to a postgres database. It was written for fun, and to get to know some of Go's API capablities.

To run this API you need to do a few things.

1) Create a database.
2) Add the database connection information as environment variables or in a `.env` file (see below).
3) Run the data.sql script against the database to create the tables and insert the default records.
3) Win!

# Installation

You need to set up your GOPATH properly, see [https://golang.org/doc/gopath_code.html](https://golang.org/doc/gopath_code.html).

Mostly, this means `$USER/go/src`.

Run this in your go path:

```
git clone https://github.com/blargism/example-go-api.git baddrivers
```

Make sure you clone to the `baddrivers` directory instead of just a straight clone. This ensures that all the go package names are correct.

# Environment Variables

The following are required to run the app, and should either be environment variables or put in a .env file.

The values in <angle brackets> should be replaced with your values.

```
PGUSER=<database user>
PGPASSWORD=<database password>
PGHOST=<database host>
PGPORT=<database port>
PGDATABASE=<database name>
```

# Go Libraries

You need to get the required libraries, run these commands:

```
go get github.com/gorilla/mux
go get github.com/gorilla/schema
go get github.com/gorilla/handlers
github.com/jackc/pgx
```

This installs some Gorilla utilities, which make handling requests easier, and PGX which is a Postgres library.
