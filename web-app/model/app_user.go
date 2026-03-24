package model

type AppUser struct {
	Id           int
	Email        string
	Role         string
	PasswordHash string
	Provider     string
	ProviderId   string
	CreatedAt    string
}
