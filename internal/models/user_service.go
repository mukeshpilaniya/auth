package models

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// GetUserByID return a user by user id
func (m *DBModel) GetUserByID(id uuid.UUID) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var u User

	row := m.DB.QueryRowContext(ctx, `
		SELECT
		       id, first_name, last_name, email, password, is_verified, created_at, updated_at
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
		       id, first_name, last_name, email, password, is_verified, created_at, updated_at
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
		&u.IsVerified,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (m *DBModel) SaveUser(user User) (User, error) {
	ctx, cancel :=context.WithTimeout(context.Background(),3*time.Second)
	defer cancel()

	stmt :=`
		insert into users
			(id, first_name, last_name, email,
			 password, is_verified, created_at, updated_at)
			values ($1,$2,$3,$4,$5,$6,$7,$8)
		`
	uid, err := uuid.NewUUID()
	if err != nil{
		return user, err
	}
	user.ID=uid
	user.IsVerified=false
	_, err = m.DB.ExecContext(ctx,stmt,
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.IsVerified,
		time.Now(),
		time.Now(),
		)
	if err != nil{
		return user, err
	}
	return user, nil
}

