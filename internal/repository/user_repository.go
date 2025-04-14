package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"user-service/internal/model"
)

// CreateUser sends a POST request to Supabase to create a new user.
func CreateUser(user *model.User) error {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	log.Println(" Supabase URL:", supabaseURL)
	log.Println(" Supabase KEY: [HIDDEN]") // Never print full key in prod

	// Construct the full Supabase endpoint URL
	url := supabaseURL + "/rest/v1/users"

	// Marshal user struct into JSON body
	body, err := json.Marshal(user)
	if err != nil {
		log.Println("Error marshaling user:", err)
		return err
	}

	// Create POST request to Supabase
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	// Add required headers for authorization
	req.Header.Set("apikey", supabaseKey)
	req.Header.Set("Authorization", "Bearer "+supabaseKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=minimal") // Optional: avoid extra response data

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return err
	}
	defer res.Body.Close()

	// Check for response errors from Supabase
	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		log.Printf("Supabase returned error: %s\n", string(bodyBytes))
		return fmt.Errorf("Failed to create user, Status: %s, Response: %s", res.Status, string(bodyBytes))
	}

	log.Println("User created successfully in Supabase")
	return nil
}

// GetUserByEmail retrieves a user by email from the Supabase database.
func GetUserByEmail(email string) (*model.User, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	url := supabaseURL + "/rest/v1/users?email=eq." + email

	// Create GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("apikey", os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_KEY"))

	// Send the request to Supabase
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Decode response JSON to user
	var users []model.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		log.Println("Error decoding response:", err)
		return nil, err
	}

	// If no user found, return error
	if len(users) == 0 {
		return nil, fmt.Errorf("User not found")
	}

	log.Println("User found:", users[0].Email)
	return &users[0], nil
}
