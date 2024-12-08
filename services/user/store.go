package user

import (
	"database/sql" // Importing the sql package for database interaction
	"errors"
	"fmt" // Importing the fmt package for formatted output and error handling

	"github.com/code-farms/go-backend/types" // Importing the custom types package for user model
)

// Store represents the storage layer for user-related operations
// It holds a reference to the database connection.
type Store struct {
	db *sql.DB  // The database connection object
}

var ErrUserNotFound = errors.New("user not found")

// NewStore creates and returns a new Store object, initialized with a database connection.
func NewStore(db *sql.DB) *Store {
	// Step 4: Initialize and return a new Store with the given database connection
	return &Store{
		db: db,  // Store the reference to the database connection
	}
}

// CreateUser is a placeholder method for creating a user.
func (s *Store) CreateUser (user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?, ?, ?, ?)", user.FirstName, user.LastName, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil  // Return nil (not implemented)
}

// GetUserByEmailId retrieves a user by email from the database and returns a User object.
// GetUserByEmail retrieves a user by their email address from the database.
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
    // Step 1: Query the database for a single user by email
    row := s.db.QueryRow("SELECT id, first_name, last_name, email, password FROM users WHERE email = ?", email)

    // Step 2: Create a new user object to hold the result
    user := new(types.User)

    // Step 3: Scan the row into the user object
    err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrUserNotFound // Return a specific error for "user not found"
        }
        return nil, err // Return any other database error
    }

    // Step 4: Return the user object and nil (no error)
    return user, nil
}

// GetUserById is a placeholder method for retrieving a user by their ID.
func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT *FROM users WHERE id = ?", id);
	if err != nil {
		return nil, err
	}

	u, err := scanRowIntoUser(rows)
	if err != nil {
		return nil, err
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return nil, nil  // Return nil for both user and error (not implemented)
}

// scanRowIntoUser is a helper function to scan a single row from the result set into a User object.
func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	// Step 12: Create a new user object to hold the scanned data
	user := new(types.User)

	// Step 13: Scan the columns from the current row into the user object
	err := rows.Scan(
		&user.ID,        // User ID
		&user.FirstName,  // User's first name
		&user.LastName,   // User's last name
		&user.Email,      // User's email address
		&user.Password,   // User's hashed password
		&user.CreatedAt,  // User's account creation date
	)

	// Step 14: If there is an error during scanning, return it
	if err != nil {
		return nil, err  // Return nil and the scanning error
	}

	// Step 15: If scanning is successful, return the user object
	return user, nil  // Return the populated user object and nil (no error)
}