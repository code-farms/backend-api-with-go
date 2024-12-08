package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/code-farms/go-backend/services/product"
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
    if s.db == nil {
        return fmt.Errorf("database connection is not initialized")
    }

    router := mux.NewRouter().StrictSlash(true)
    subRouter := router.PathPrefix("/api/v1/").Subrouter()

    userStore := user.NewStore(s.db)
    userHandler := user.NewHandler(userStore)
    userHandler.RegisterRoutes(subRouter)

    productStore := product.NewStore(s.db)
    productHandler := product.NewHandler(productStore)
    productHandler.RegisterRoutes(subRouter)

    log.Printf("Server is starting on %s...", s.addr)
    err := http.ListenAndServe(s.addr, router)
    if err != nil {
        log.Printf("Error starting server: %v", err)
    }

    return err
}