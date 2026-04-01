package controller

import (
	"algtutor/domain"
	"algtutor/service"
	"encoding/json"
	"net/http"
)

type AppUserController struct {
	aus *service.AppUserService
}

func NewAppUserController(aus *service.AppUserService) *AppUserController {
	return &AppUserController{aus}
}

func (auc *AppUserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email        string `json:"email"`
		Role         string `json:"role"`
		PasswordHash string `json:"passwordHash"`
		Provider     string `json:"provider"`
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	au := domain.AppUser{
		Email:        req.Email,
		Role:         req.Role,
		PasswordHash: req.PasswordHash,
		Provider:     req.Provider,
	}
	err = auc.aus.CreateUser(au)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
