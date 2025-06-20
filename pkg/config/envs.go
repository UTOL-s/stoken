// Package config provides functionality for accessing configuration settings
// and environment variables required by the application.
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// SuperTokenURL retrieves the SuperTokens service URL from environment variables.
// It loads environment variables from a .env file using godotenv.
// If the .env file cannot be loaded, the function will log a fatal error and terminate the program.
// Returns the value of the SUPER_TOKENS_URL environment variable.
func SuperTokenURL() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("SUPER_TOKENS_URL ERROR")
	}

	return os.Getenv("SUPER_TOKENS_URL")
}

// SuperTokensKey retrieves the SuperTokens API key from environment variables.
// It loads environment variables from a .env file using godotenv.
// If the .env file cannot be loaded, the function will log a fatal error and terminate the program.
// Returns the value of the SUPER_TOKENS_KEY environment variable.
func SuperTokensKey() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("SUPER_TOKENS_KEY ERROR")
	}

	return os.Getenv("SUPER_TOKENS_KEY")
}

// EmailHost retrieves the email server host address from environment variables.
// It loads environment variables from a .env file using godotenv.
// If the .env file cannot be loaded, the function will log a fatal error and terminate the program.
// Returns the value of the EMAIL_HOST environment variable.
func EmailHost() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("EMAIL_HOST ERROR")
	}

	return os.Getenv("EMAIL_HOST")
}

// EmailFrom retrieves the sender email address from environment variables.
// It loads environment variables from a .env file using godotenv.
// If the .env file cannot be loaded, the function will log a fatal error and terminate the program.
// Returns the value of the EMAIL_FROM environment variable.
func EmailFrom() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("EMAIL_FROM ERROR")
	}

	return os.Getenv("EMAIL_FROM")
}

// EmailPassword retrieves the email account password from environment variables.
// It loads environment variables from a .env file using godotenv.
// If the .env file cannot be loaded, the function will log a fatal error and terminate the program.
// Returns the value of the EMAIL_PASSWORD environment variable.
func EmailPassword() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("EMAIL_PASSWORD ERROR")
	}

	return os.Getenv("EMAIL_PASSWORD")
}
