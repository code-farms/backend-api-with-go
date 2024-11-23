package auth

import "golang.org/x/crypto/bcrypt" // Import the bcrypt package for hashing passwords

// HashPassword takes a plain-text password and returns the hashed version of it
func HashPassword(password string) (string, error) {
	// Step 1: Hash the password using bcrypt with the default cost factor
	// `bcrypt.GenerateFromPassword()` generates a bcrypt hash from the given password
	// The `bcrypt.DefaultCost` is a predefined cost factor that determines the complexity of the hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// Step 2: If there is an error while hashing the password, return an empty string and the error
	if err != nil {
		return "", err  // Return an error if hashing fails
	}

	// Step 3: Return the hashed password as a string (since `GenerateFromPassword` returns a byte slice)
	return string(hash), nil  // Return the hashed password as a string and nil error on success
}