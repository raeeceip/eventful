package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

// jwt generates a JWT secret key and writes it to the .env file.
func jwt() {
	length := 32
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	secret := base64.URLEncoding.EncodeToString(bytes)
	fmt.Println("Generated JWT_SECRET:", secret)

	// Define the .env file path
	envFilePath := filepath.Join(".", ".env")

	// Open the .env file, create it if it doesn't exist
	file, err := os.OpenFile(envFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Write the JWT secret to the .env file
	_, err = file.WriteString(fmt.Sprintf("JWT_SECRET=%s\n", secret))
	if err != nil {
		panic(err)
	}

	fmt.Println("JWT_SECRET written to .env file")
}
