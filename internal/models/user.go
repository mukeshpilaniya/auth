package models

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// User is a type for all users
type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

// GetUserByID return a user by user id
func (m *DBModel) GetUserByID(id int64) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u User

	row := m.DB.QueryRowContext(ctx, `
		SELECT
		       id, first_name, last_name, email, password
		FROM 
		     users
		WHERE
			id = $1`, id)

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}

// GetUserByEmail return a user by email address
func (m *DBModel) GetUserByEmail(email string) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u User

	row := m.DB.QueryRowContext(ctx, `
		SELECT
		       id, first_name, last_name, email, password
		FROM 
		     users
		WHERE
			email = $1`, email)

	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}

//func (m *DBModel) saveUser() (User, error) {
//	ctx, cancel :=context.WithTimeout(context.Background(),3*time.Second)
//	defer cancel()
//
//	var u User
//
//	row := m.DB.QueryRowContext(ctx,`
//		SELECT
//		       id, user_name, first_name, last_name, password
//		FROM
//		     users
//		WHERE
//			$1`, id)
//
//	err := row.Scan(
//		&u.ID,
//		&u.FirstName,
//		&u.LastName,
//		&u.UserName,
//		&u.Password,
//	)
//	if err !=nil {
//		return u, err
//	}
//	return u, nil
//}
