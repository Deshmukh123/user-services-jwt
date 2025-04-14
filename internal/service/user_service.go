package service

import (
	"log"
	"user-service/internal/model"
	"user-service/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// Register hashes the user's password before saving to the database.
func Register(user *model.User) error {
	// Hash the plain password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(" Error hashing password:", err)
		return err
	}

	user.Password = string(hash) // Save hashed password
	log.Println("Hashing password before sending to Supabase:", user.Password)

	err = repository.CreateUser(user)
	if err != nil {
		log.Println(" Error creating user:", err)
		return err
	}

	log.Println(" User registered:", user.Email)
	return nil
}

// Login checks credentials and returns the user if valid.
func Login(email, password string) (*model.User, error) {
	// Retrieve the user from the database
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		log.Println(" User not found:", err)
		return nil, err
	}

	log.Println("Checking password for:", email)

	// Compare the hashed password from DB with the plain password input
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(" Error hashing password:", user.Password, err)

		log.Println(" Incorrect password:", err)
		return nil, err
	}

	log.Println(" Login successful for:", email)
	return user, nil
}
