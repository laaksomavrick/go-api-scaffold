package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// User defines the shape of a User for our application
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) validateForCreate() error {
	if u.Email == "" || u.Password == "" {
		return errors.New("user: missing required values")
	}
	if len(u.Email) < 8 || len(u.Password) < 8 {
		return errors.New("user: email and password must be greater than 8 characters in length")
	}
	return nil
}

func (u *User) prepareForInsert() error {
	password := []byte(u.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) compareHashAndPassword(password string) error {
	p := []byte(password)
	hp := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(hp, p)
}
