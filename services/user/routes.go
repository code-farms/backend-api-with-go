package user

import (
	"errors"
	"fmt"      // Importing fmt for formatted output (error messages)
	"net/http" // Importing net/http for handling HTTP requests and responses

	"github.com/code-farms/go-backend/configs"
	"github.com/code-farms/go-backend/services/auth" // Importing the auth package for password hashing
	"github.com/code-farms/go-backend/types"         // Importing the custom types for user and payload definitions
	"github.com/code-farms/go-backend/utils"         // Importing utility functions for parsing and writing JSON
	"github.com/go-playground/validator/v10"         // Importing the validator package for data validation
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
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Register route for handling login, which will invoke the handleLogin method for POST requests
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	// Register route for handling user registration, which will invoke the handleRegister method for POST requests
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// handleLogin is the placeholder function for the login route.
// It receives the request, validates the input, checks if the user exists, and generates a JWT token.
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
    // Step 1: Parse the request body into the LoginUserPayload struct.
    var payload types.LoginUserPayload
    if err := utils.ParseJSON(r, &payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
        return
    }

    // Step 2: Validate the payload using the validator
    if err := utils.Validate.Struct(payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", err))
        return
    }

    // Step 3: Retrieve the user by email
    user, err := h.store.GetUserByEmail(payload.Email)
    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            // If no user is found, return 401 Unauthorized
            utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
            return
        }
        // For other unexpected errors, return 500 Internal Server Error
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch user: %v", err))
        return
    }

    // Step 4: Compare the hashed password
    if !auth.ComparePasswords(user.Password, []byte(payload.Password)) {
        utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid email or password"))
        return
    }

    // Step 5: Generate a JWT token
    secret := []byte(configs.Envs.JWTSecret) // Use handler-level configuration for flexibility
    token, err := auth.CreateJWT(secret, user.ID)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create JWT token"))
        return
    }

    // Step 6: Return the JWT token in the response
    utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

// handleRegister is the function that handles user registration requests.
// It receives the request, validates the input, checks if the user exists, and stores the user in the database.
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Step 1: Parse the request body into the RegisterUserPayload struct.
	var payload types.RegisterUserPayload
    err := utils.ParseJSON(r, &payload)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON: %v", err))
        return
    }

	 // Step 2: Validate the payload using the validator
	 if err := utils.Validate.Struct(payload); err != nil {
        errors := err.(validator.ValidationErrors)
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
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
    if err != nil || hashedPassword == "" {
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to hash password"))
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
        utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create user: %v", err))
        return
    }

	// Step 9: Return a 201 Created response
    utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user registered successfully"})
}