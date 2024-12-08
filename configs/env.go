package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv" // Importing the package to load environment variables from .env file
)

// Config structure defines the configuration variables for the application
type Config struct {
	PublicHost string  // URL where the public-facing service will be hosted
	Port       string  // Port for the server to listen on
	DBUser     string  // Database username
	DBPassword string  // Database password
	DBAddress  string  // Database host and port, formatted as "host:port"
	DBName     string  // Name of the database
	JWTExpirationInSeconds int64 // JWT expiration time
	JWTSecret string // JWT secret key
}

// Envs variable holds the application configuration, initialized using initConfig()
var Envs = initConfig()  // Global variable initialized with the result of initConfig()

// initConfig loads environment variables and initializes the Config structure
func initConfig() Config {
	// Step 1: Load the environment variables from the .env file (if present)
	// godotenv.Load() will load the variables from a `.env` file into the application's environment
	godotenv.Load()

	// Step 2: Return a Config structure with values loaded from environment variables
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),  // Default: "http://localhost"
		Port: getEnv("PORT", "8080"),  // Default: "8080"
		DBUser: getEnv("DB_USER", "root"),  // Default: "root"
		DBPassword: getEnv("DB_PASSWORD", "root"),  // Default: "root"
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),  // Default: "127.0.0.1:3306"
		DBName: getEnv("DB_NAME", "go_backend"),  // Default: "go_backend"
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION", 3600 * 24 * 7),  // Default: 3600 seconds (1 hour)
	}
}

// getEnv is a helper function that retrieves an environment variable value
// If the value is not found, it returns the fallback value.
func getEnv(key, fallback string) string {
	// Step 9: Lookup the environment variable using os.LookupEnv()
	if value, ok := os.LookupEnv(key); ok {
		return value  // If the key exists in the environment, return the value
	}
	// Step 10: If the key does not exist, return the fallback value
	return fallback  // Return the fallback value if the environment variable is not set
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64);
		if  err != nil {
			return fallback
		}
		return i
	}
	return fallback
}