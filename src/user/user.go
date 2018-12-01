package user

import "errors"

// User defines the shape of a User for our application
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) validateForCreate() error {
	if u.Email == "" || u.Password == "" {
		return errors.New("user: found nil values")
	}
	if len(u.Email) < 8 || len(u.Password) < 8 {
		return errors.New("user: email and password must be greater than 8 characters in length")
	}
	return nil
}
