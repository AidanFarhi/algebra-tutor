package domain

import "time"

type Session struct {
	Id        int
	UserId    int
	CreatedAt time.Time
	ExpiresAt time.Time
}

func NewSession(id int, userId int, createdAt time.Time, expiresAt time.Time) (Session, error) {
	// TODO: some validations maybe?
	return Session{id, userId, createdAt, expiresAt}, nil
}
