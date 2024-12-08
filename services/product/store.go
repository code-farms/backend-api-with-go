package product

import (
	"database/sql" // Importing the sql package for database interaction
	"fmt"

	"github.com/code-farms/go-backend/types"
)

type store struct {
	db *sql.DB // Declaring a variable 'db' of type *sql.DB
}	

func NewStore (db *sql.DB) *store {
	return &store{
		db: db, // Assigning the provided database connection to the 'db' field
	}
}

func (s *store) CreateProduct(product types.CreateProductPayload) (int64, error) {
    // Use a parameterized query to prevent SQL injection
    query := "INSERT INTO products (name, price, image, description, quantity) VALUES (?, ?, ?, ?, ?)"
    
    // Execute the query and capture the result
    result, err := s.db.Exec(query, product.Name, product.Price, product.Image, product.Description, product.Quantity)
    if err != nil {
        return 0, fmt.Errorf("failed to insert product into database: %w", err)
    }

    // Fetch the ID of the inserted row (if needed)
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("failed to fetch inserted product ID: %w", err)
    }

    // Return the inserted product ID and nil error
    return id, nil
}


func (s *store) GetProducts () ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products") // Querying the database to retrieve all products
	if err != nil {
		return nil, err // Returning an error if the query fails
	}
	
	products := make([]*types.Product, 0) // Creating a slice to store the retrieved products
	for rows.Next() {
		p, err := scanRowIntoProduct(rows) // Scanning each row into a Product struct
		if err != nil {
			return nil, err // Returning an error if scanning fails
		}

		products = append(products, p) // Adding the scanned Product to the slice
	}

	return products, nil // Returning the slice of products and a nil error
}

// func (s *store) GetProductById (id int) (*Product, error) {
// 	return nil, nil // Placeholder implementation, to be implemented later
// }


func scanRowIntoProduct (rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)
	err := rows.Scan(
		&product.ID,       // Product ID
		&product.Name,     // Product name
		&product.Description, // Product description
		&product.Image,    // Product image
		&product.Quantity, // Product quantity
		&product.Price,    // Product price
		&product.CreatedAt, // Product creation date
	)
	return product, err
}