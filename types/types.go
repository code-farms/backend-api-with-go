package types

import (
	"time"
)

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

type ProductStore interface {
	GetProductByID(id int) (*Product, error)
	GetProductsByID(ids []int) ([]Product, error)
	GetProducts() ([]*Product, error)
	CreateProduct(CreateProductPayload) error
	UpdateProduct(Product) error
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

// Product represents a product in the system.	
// It contains all the necessary fields required to store product information in the database.
type Product struct {
	ID        int       `json:"id"`        // The unique identifier for the product
	Name      string    `json:"name"`      // The name of the product
	Description string    `json:"description"`  // The description of the product
	Image     string    `json:"image"`     // The image URL of the product
	Quantity  int       `json:"quantity"`  // The quantity of the product
	Price     float64   `json:"price"`     // The price of the product
	CreatedAt time.Time `json:"createdAt"`  // The timestamp when the product was created in the system
}

// RegisterUserPayload represents the data required to register a new user.
// This is the structure that the client will send in the request body when registering.
type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`  // The first name of the user
	LastName  string `json:"lastName" validate:"required"`   // The last name of the user
	Email     string `json:"email" validate:"required,email"`      // The email address of the user (unique)
	Password  string `json:"password" validate:"required,min=3,max=130"`   // The password of the user (to be hashed before storing)
}

// LoginUserPayload represents the data required to log in a user.
// This is the structure that the client will send in the request body when logging in.
type LoginUserPayload struct {
	Email     string `json:"email" validate:"required,email"`      // The email address of the user (unique)
	Password  string `json:"password" validate:"required"`   // The password of the user (to be hashed before storing)
}

// CreateProductPayload represents the data required to create a new product.
// This is the structure that the client will send in the request body when creating a new product.
type CreateProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}