package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http" // ye import bhul gaya tu
	"os"
	"user-service/internal/db"
	"user-service/internal/model"
)

func CreateUser(user *model.User) error {
	log.Println("SupabaseURL:", db.SupabaseURL)
	log.Println("SupabaseKey:", os.Getenv("SUPABASE_KEY"))

	log.Println("Supabase", os.Getenv("SUPABASE_URL"))
	log.Println("Supabase", os.Getenv("SUPABASE_KEY"))
	url := os.Getenv("SUPABASE_URL") + "/rest/v1/users"

	body, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("apikey", os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		return fmt.Errorf("Failed to create user, Status: %s, Response: %s", res.Status, string(bodyBytes))
	}

	return nil

	// req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	// req.Header.Set("apikey", os.Getenv("SUPABASE_KEY"))
	// req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_KEY"))
	// req.Header.Set("Content-Type", "application/json")

	// resp, err := db.HttpClient.Do(req)
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode >= 400 {
	// 	return fmt.Errorf("Failed to create user, Status: %s", resp.Status)
	// }
	// return nil
}

func GetUserByEmail(email string) (*model.User, error) {
	url := db.SupabaseURL + "/rest/v1/users?email=eq." + email

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", os.Getenv("SUPABASE_KEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SUPABASE_KEY"))

	resp, err := db.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []model.User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("User not found")
	}

	return &users[0], nil
}
