package domain

import "errors"

type AppUser struct {
	Id           int
	Email        string
	Role         string
	PasswordHash string
	Provider     string
	ProviderId   string
	CreatedAt    string
}

func (u AppUser) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Role == "" {
		return errors.New("role is required")
	}
	if u.PasswordHash == "" {
		return errors.New("password is required")
	}
	if u.Provider == "" {
		return errors.New("provider is required")
	}
	return nil
}
