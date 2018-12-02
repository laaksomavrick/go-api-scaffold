package user

import "database/sql"

type userRepository struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) insert(user *User) error {
	query := `
			INSERT INTO users (email, password)
			VALUES ($1, $2)
			RETURNING id`

	err := ur.db.QueryRow(query, user.Email, user.Password).Scan(&user.ID)
	return err
}
