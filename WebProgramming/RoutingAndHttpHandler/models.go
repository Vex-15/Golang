package main

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(name, email, hashedPassword, avatar string) (int64, error)
	GetUserByEmail(email string) (*User, error)
	GetUsers() ([]User, error)
}

type SQLUserRepository struct {
	db *sql.DB
}

// NewSQLUserRepository creates a new repository
func NewSQLUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{
		db: db,
	}
}

func (r *SQLUserRepository) CreateUser(name, email, hashedPassword, avatar string) (int64, error) {
	ctx := context.Background()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Hash password
	hp, err := bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// Insert user and get generated ID
	var userID int64

	err = tx.QueryRowContext(
		ctx,
		`INSERT INTO users (name, email, hashed_password)
		 VALUES ($1, $2, $3)
		 RETURNING id`,
		name,
		email,
		string(hp),
	).Scan(&userID)

	if err != nil {
		return 0, err
	}

	// Insert profile
	_, err = tx.ExecContext(
		ctx,
		`INSERT INTO profile (user_id, avatar)
		 VALUES ($1, $2)`,
		userID,
		avatar,
	)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *SQLUserRepository) GetUserByEmail(email string) (*User, error) {

	stmt := `
	SELECT
		u.id,
		u.name,
		u.email,
		u.hashed_password,
		u.created_at,
		p.avatar
	FROM users u
	INNER JOIN profile p
		ON u.id = p.user_id
	WHERE u.email = $1
	`

	row := r.db.QueryRow(stmt, email)

	var user User

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.Profile.Avatar,
	)
	if err != nil {
		return nil, err
	}

	user.Profile.UserID = user.ID

	return &user, nil
}

func (r *SQLUserRepository) GetUsers() ([]User, error) {

	stmt := `
	SELECT
		id,
		name,
		email,
		hashed_password,
		created_at
	FROM users
	`

	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.HashedPassword,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
