package user

import (
	"fmt"      // Importing fmt for formatted output (error messages)
	"net/http" // Importing net/http for handling HTTP requests and responses

	"github.com/code-farms/go-backend/services/auth" // Importing the auth package for password hashing
	"github.com/code-farms/go-backend/types"         // Importing the custom types for user and payload definitions
	"github.com/code-farms/go-backend/utils"         // Importing utility functions for parsing and writing JSON
	"github.com/gorilla/mux"                         // Importing the gorilla/mux router package for handling routes
)

// Handler struct holds the reference to the UserStore interface
// which will be used to interact with the database.
type Handler struct {
	store types.UserStore  // A reference to the UserStore interface for interacting with user data
}

// NewHandler is a constructor function that creates and returns a new Handler object
// initialized with a store for user data interaction.
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}  // Return a new Handler with the store
}

// ResisterRoutes method registers the routes for login and register with the provided router.
// It maps HTTP methods (POST) to their respective handler functions.
func (h *Handler) ResisterRoutes(router *mux.Router) {
	// Register route for handling login, which will invoke the handleLogin method for POST requests
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	// Register route for handling user registration, which will invoke the handleRegister method for POST requests
	router.HandleFunc("/resister", h.handleRegister).Methods("POST")
}

// handleLogin is the placeholder function for the login route.
// Currently, it's unimplemented, but it will handle user login requests in the future.
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// This is left unimplemented for now
}

// handleRegister is the function that handles user registration requests.
// It receives the request, validates the input, checks if the user exists, and stores the user in the database.
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Step 1: Parse the request body into the RegisterUserPayload struct.
	var payload types.RegisterUserPayload
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		// Step 2: If the parsing fails, return a Bad Request (400) error with the parsing error
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Step 3: Check if a user already exists with the provided email.
	_, err = h.store.GetUserByEmail(payload.Email)
	if err == nil {
		// Step 4: If a user with the same email already exists, return a conflict (400) error
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// Step 5: Hash the user's password using the `auth.HashPassword` function
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		// Step 6: If password hashing fails, return an internal server error (500)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Step 7: Create a new User object with the parsed data and hashed password
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,  // First name from the payload
		LastName: payload.LastName,    // Last name from the payload
		Email: payload.Email,          // Email from the payload
		Password: hashedPassword,      // Hashed password
	})

	if err != nil {
		// Step 8: If user creation fails, return an internal server error (500)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Step 9: If user creation is successful, return a 201 Created status with no content
	utils.WriteJSON(w, http.StatusCreated, nil)
}