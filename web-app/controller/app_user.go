package controller

import (
	"algtutor/service"
	"fmt"
	"net/http"
)

type AppUserController struct {
	aus *service.AppUserService
}

func NewAppUserController(aus *service.AppUserService) *AppUserController {
	return &AppUserController{aus}
}

func (auc *AppUserController) Register(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("register-email")
	password := r.FormValue("register-password")
	passwordRepeat := r.FormValue("register-password-repeat")
	role := "user"
	provider := "local"
	session, err := auc.aus.RegisterAndLogin(email, password, passwordRepeat, role, provider)
	if err != nil {
		// figure out what the error is and then render accordingly
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(session.Id, session.UserId)
	// set the cookie
	// cookie := &http.Cookie{
	// 	Name:     "session_id",
	// 	Value:    session.Id,
	// 	Path:     "/",
	// 	HttpOnly: true,
	// 	Secure:   true,
	// 	SameSite: http.SameSiteStrictMode,
	// 	MaxAge:   24 * 60 * 60 * 30, // 30 days in seconds
	// }
	// http.SetCookie(w, cookie)
	// http.Redirect(w, r, "/", http.StatusFound)
}
