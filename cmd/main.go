package main

import (
	"database/sql"
	"log"

	"github.com/code-farms/go-backend/cmd/api" // Importing the API server package
	"github.com/code-farms/go-backend/configs" // Importing configurations (environment variables)
	"github.com/code-farms/go-backend/db"      // Importing the database handling package
	"github.com/go-sql-driver/mysql"           // Importing the MySQL driver
)

func main() {
	// Step 1: Set up MySQL database configuration
	// Step 2: Connect to the MySQL database
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 configs.Envs.DBUser,            // MySQL username
		Passwd:               configs.Envs.DBPassword,        // MySQL password
		Addr:                 configs.Envs.DBAddress,         // MySQL server address
		DBName:               configs.Envs.DBName,            // Database name
		Net:                  "tcp",                          // Connection protocol
		AllowNativePasswords: true,                           // Allow native passwords
		ParseTime:            true,                           // Parse time values correctly
	})

	// Step 3: Handle database connection error
	if err != nil {
		log.Fatal(err) // Log the error and stop the application
	}

	// Step 4: Check the database connectivity
	// Step 5: Handle connection error
	// Step 6: Log successful connection
	intiStorage(db)

	// Step 7: Create a new API server
	// Step 8: Run the server
	// Step 9: Handle server error
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err) // Log the error and stop the application
	}
}

// intiStorage checks the database connectivity and logs the status
func intiStorage(db *sql.DB) {
	// Step 4: Ping the database to ensure the connection is valid
	err := db.Ping()
	// Step 5: Handle connection error
	if err != nil {
		log.Fatal(err) // Log the error and stop the application if the database is unreachable
	}

	// Step 6: Log a success message if the database is connected successfully
	log.Println("Connected to MySQL database")
}