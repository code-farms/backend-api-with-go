package db

import (
	"database/sql" // Importing the SQL package for database interaction
	"log"          // Importing the log package for error logging

	"github.com/go-sql-driver/mysql" // Importing the MySQL driver package to interact with MySQL database
)

// NewMySQLStorage initializes a new MySQL database connection using the provided configuration
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	// Step 1: Open a connection to the MySQL database using the provided configuration
	// `sql.Open()` takes the driver name ("mysql") and a data source name (DSN) formatted by `cfg.FormatDSN()`
	db, err := sql.Open("mysql", cfg.FormatDSN())

	// Step 2: If there is an error while opening the connection, log the error and terminate the application
	if err != nil {
		log.Fatal(err)  // Log the error and stop execution if the connection cannot be established
	}

	// Step 3: Return the database connection object and a nil error indicating successful initialization
	return db, nil
}