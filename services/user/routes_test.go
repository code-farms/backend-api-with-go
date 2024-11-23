package user

import (
	"bytes"             // For creating a buffer for the request body
	"encoding/json"     // For marshaling the payload into JSON format
	"fmt"               // For formatted I/O, used to create error messages
	"net/http"          // For HTTP handling functions like NewRequest, MethodPost, etc.
	"net/http/httptest" // For creating a test HTTP server and recording responses
	"testing"           // For writing unit tests

	"github.com/code-farms/go-backend/types" // Importing the types package for the user payload and user structure
	"github.com/gorilla/mux"                 // Importing the Gorilla mux router for routing HTTP requests
)

// TestUserServiceHandlers is the test function that will validate the user service handlers.
func TestUserServiceHandlers(t *testing.T) {
	// Mock the UserStore interface to simulate the behavior of the data store during testing.
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)  // Create a new handler with the mock user store

	// Test Case 1: Test if the handler fails when the payload is invalid.
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		// Create an invalid payload with incorrect email and other invalid data.
		payload := types.RegisterUserPayload{
			FirstName: "John",  // Valid first name
			LastName: "123",  // Invalid last name (should be a string, this is a number)
			Password: "adsdf",  // Valid password (not hashed)
			Email: "dfh",  // Invalid email (does not match standard email format)
		}

		// Marshal the payload into JSON format to send in the HTTP request body
		marshalled, _ := json.Marshal(payload)

		// Create a new POST request with the JSON payload
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))

		if err != nil {
			// If there's an error creating the request, fail the test
			t.Fatal(err)
		}

		// Create a new HTTP test recorder to record the response from the handler
		rr := httptest.NewRecorder()

		// Set up the router and register the handler
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)  // Register the route

		// Serve the HTTP request and capture the response in the recorder
		router.ServeHTTP(rr, req)

		// Assert that the response status code is 400 (Bad Request) for invalid input
		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d but got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

// mockUserStore is a mock implementation of the UserStore interface used for testing purposes.
type mockUserStore struct{}

// GetUserByEmail simulates the behavior of fetching a user by email.
// It returns an error indicating that the user is not found, as this is a mock implementation.
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

// GetUserById simulates the behavior of fetching a user by ID.
// Since it's a mock, it simply returns nil.
func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

// CreateUser simulates the behavior of creating a user in the database.
// It returns nil (no error) to indicate successful user creation in this mock.
func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}