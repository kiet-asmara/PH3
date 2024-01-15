package models

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   uint   `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserService struct {
	DB *sql.DB
}

// use pointer in one func = use it everywhere (can return nil)
func (us *UserService) Create(input UserInput) (*User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	passwordHash := string(hashedPass)

	user := User{
		Email:    input.Email,
		Password: passwordHash,
	}

	row := us.DB.QueryRow(`INSERT INTO users (email, password) VALUES ($1, $2) RETURNING user_id`, input.Email, passwordHash)
	err = row.Scan(&user.UserID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

func (us *UserService) Read(input UserInput) (*User, error) {
	row := us.DB.QueryRow(`SELECT * FROM users WHERE email = $1`, input.Email)

	var user User
	err := row.Scan(&user.UserID, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("read user: %w", err)
	}

	return &user, nil
}
