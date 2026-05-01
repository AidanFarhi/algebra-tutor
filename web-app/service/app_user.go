package service

import (
	"algtutor/domain"
	"algtutor/repo"
	"time"
)

type AppUserService struct {
	aur *repo.AppUserRepo
}

func NewAppUserService(aur *repo.AppUserRepo) *AppUserService {
	return &AppUserService{aur}
}

func (aus AppUserService) RegisterAndLogin(email, password, passwordRepeat, role, provider string) (domain.Session, error) {

	// check if user exists by email

	// hash password

	// create user domain object

	// pass to repo to create

	// create a session

	// return session info to controller

	return domain.NewSession(1, 1, time.Time{}, time.Time{})
}

func (aus AppUserService) GetUser(id int) (domain.AppUser, error) {
	return aus.aur.GetByID(id)
}
