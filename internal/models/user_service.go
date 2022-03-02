package models

import (
	"context"
	"time"
)

// GetUserByID return a user by user id
func (m *DBModel) GetUserByID(id int64) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u User

	row := m.DB.QueryRowContext(ctx, `
		SELECT
		       id, first_name, last_name, email, password, token_hash, is_verified, created_at, updated_at
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
		&u.TokenHash,
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
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
		       id, first_name, last_name, email, password, token_hash, is_verified, created_at, updated_at
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
		&u.TokenHash,
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
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

