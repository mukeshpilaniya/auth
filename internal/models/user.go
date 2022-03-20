package models

import (
	"github.com/google/uuid"
	"time"
)

// User is a type for all users
type User struct {
	ID         uuid.UUID `json:"id,omitempty"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	IsVerified bool      `json:"is_verified"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// NewUser return an instance of user
func NewUser(firstName string, lastName string, email string, password string)(*User, error){
	id, err := uuid.NewUUID()
	if err != nil{
		return &User{}, err
	}
	return &User{
		ID: id,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		Password: password,
		IsVerified: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}