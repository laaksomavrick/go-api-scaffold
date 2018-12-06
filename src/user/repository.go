package user

import (
	"database/sql"
)

// Repository wraps common DB operations for users
type Repository struct {
	db *sql.DB
}

// NewRepository constructs a Repository
func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// Insert inserts a record into the users table
func (ur *Repository) Insert(user *User) error {
	query := `
			INSERT INTO users (email, password)
			VALUES ($1, $2)
			RETURNING id`

	err := ur.db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	return err
}

// FindByEmail finds a user by email if they exist
func (ur *Repository) FindByEmail(email string) (User, error) {
	var user User
	query := "SELECT * FROM users WHERE email = $1"
	err := ur.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}

// FindByEmail finds a user by email if they exist
func (ur *Repository) FindById(id int) (User, error) {
	var user User
	query := "SELECT * FROM users WHERE id = $1"
	err := ur.db.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password)
	return user, err
}
