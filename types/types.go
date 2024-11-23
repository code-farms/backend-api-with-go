package types

import "time"

// UserStore is an interface that defines the methods required to interact with the user data.
// These methods are intended for storing, retrieving, and managing user data in the database.
type UserStore interface {
	// GetUserByEmail retrieves a user by their email address.
	// Returns a pointer to the User object and an error if the user is not found.
	GetUserByEmail(email string) (*User, error)

	// GetUserById retrieves a user by their unique ID.
	// Returns a pointer to the User object and an error if the user is not found.
	GetUserById(id int) (*User, error)

	// CreateUser stores a new user in the database.
	// Returns an error if there is an issue with storing the user data.
	CreateUser(User) error
}

// User represents a user in the system.
// It contains all the necessary fields required to store user information in the database.
type User struct {
	ID        int       `json:"id"`        // The unique identifier for the user
	FirstName string    `json:"firstName"`  // The user's first name
	LastName  string    `json:"lastName"`   // The user's last name
	Email     string    `json:"email"`      // The user's email address (unique)
	Password  string    `json:"-"`          // The user's password (never returned in the JSON response)
	CreatedAt time.Time `json:"createdAt"`  // The timestamp when the user was created in the system
}

// RegisterUserPayload represents the data required to register a new user.
// This is the structure that the client will send in the request body when registering.
type RegisterUserPayload struct {
	FirstName string `json:"firstName"`  // The first name of the user
	LastName  string `json:"lastName"`   // The last name of the user
	Email     string `json:"email"`      // The email address of the user (unique)
	Password  string `json:"password"`   // The password of the user (to be hashed before storing)
}