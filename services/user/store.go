package user

import (
	"database/sql" // Importing the sql package for database interaction
	"fmt"          // Importing the fmt package for formatted output and error handling

	"github.com/code-farms/go-backend/types" // Importing the custom types package for user model
)

// Store represents the storage layer for user-related operations
// It holds a reference to the database connection.
type Store struct {
	db *sql.DB  // The database connection object
}

// CreateUser implements the UserStore interface to create a new user in the database.
func (s *Store) CreateUser(user types.User) error {
	// Step 1: Here, we're simply panicking since the method is unimplemented
	panic("unimplemented")  // The method to create a user is not yet implemented
}

// GetUserByEmail implements the UserStore interface to fetch a user by their email.
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	// Step 2: Here, we're simply panicking since the method is unimplemented
	panic("unimplemented")  // The method to get a user by email is not yet implemented
}

// GetUserById implements the UserStore interface to fetch a user by their ID.
func (s *Store) GetUserById(id int) (*types.User, error) {
	// Step 3: Here, we're simply panicking since the method is unimplemented
	panic("unimplemented")  // The method to get a user by ID is not yet implemented
}

// NewStore creates and returns a new Store object, initialized with a database connection.
func NewStore(db *sql.DB) *Store {
	// Step 4: Initialize and return a new Store with the given database connection
	return &Store{
		db: db,  // Store the reference to the database connection
	}
}

// GetUserByEmailId retrieves a user by email from the database and returns a User object.
func (s *Store) GetUserByEmailId(email string) (*types.User, error) {
	// Step 5: Query the database to fetch a user by their email
	rows, err := s.db.Query("SELECT *FROM users WHERE email = ?", email)
	if err != nil {
		// Step 6: If an error occurs while querying the database, return it
		return nil, err  // Return nil and the error
	}

	// Step 7: Create a new empty user object to store the result
	u := new(types.User)
	for rows.Next() {
		// Step 8: Scan the result row into the user object
		u, err = scanRowIntoUser(rows)
		if err != nil {
			// Step 9: If there is an error scanning the row, return it
			return nil, err  // Return nil and the error
		}
	}

	// Step 10: If the user ID is 0, it means no user was found, so return an error
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")  // Return a "user not found" error
	}

	// Step 11: If the user is found, return the user object and nil (no error)
	return u, nil  // Return the user object and nil (no error)
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

// GetUserById is a placeholder method for retrieving a user by their ID.
func GetUserById(id int) (*types.User, error) {
	// Step 16: This method is not yet implemented, so we return nil
	return nil, nil  // Return nil for both user and error (not implemented)
}

// CreateUser is a placeholder method for creating a user.
func CreateUser(user *types.User) error {
	// Step 17: This method is not yet implemented, so we return nil
	return nil  // Return nil (not implemented)
}