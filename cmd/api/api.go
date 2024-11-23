package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/code-farms/go-backend/services/user" // Import the user service package
	"github.com/gorilla/mux"                         // Import Gorilla Mux for routing
)

// APIServer structure holds the address and database connection for the server
type APIServer struct {
	addr string  // Address to bind the server (e.g., ":8080")
	db   *sql.DB // Database connection used for API operations
}

// NewAPIServer initializes and returns a new APIServer instance
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	// Step 1: Create and return an APIServer instance with the specified address and database connection
	return &APIServer{
		addr: addr,  // Set server address
		db:   db,    // Set database connection
	}
}

// Run starts the server and listens for incoming HTTP requests
func (s *APIServer) Run() error {
	// Step 2: Create a new router using Gorilla Mux
	router := mux.NewRouter()

	// Step 3: Create a sub-router that handles routes prefixed with "/api/v1"
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// Step 4: Initialize the user store with the provided database connection
	userStore := user.NewStore(s.db)

	// Step 5: Create a new user handler that will manage user-related API operations
	userHandler := user.NewHandler(userStore)

	// Step 6: Register the user-related routes with the sub-router
	userHandler.ResisterRoutes(subRouter)

	// Step 7: Log the message indicating that the server is listening on the specified port
	log.Println("Listening on port", s.addr)

	// Step 8: Start the HTTP server to listen for incoming requests on the specified address and router
	//         This method will block and continue running until the server is stopped
	return http.ListenAndServe(s.addr, router)
}