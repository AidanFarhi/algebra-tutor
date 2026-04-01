package service

import (
	"algtutor/domain"
	"algtutor/repo"
)

type AppUserService struct {
	aur *repo.AppUserRepo
}

func NewAppUserService(aur *repo.AppUserRepo) *AppUserService {
	return &AppUserService{aur}
}

func (aus AppUserService) CreateUser(au domain.AppUser) error {
	if err := au.Validate(); err != nil {
		return err
	}
	return aus.aur.CreateUser(au)
}

func (aus AppUserService) GetUser(id int) (domain.AppUser, error) {
	return aus.aur.GetByID(id)
}
