package services

import (
	dbp "dev-blog/db"
	"dev-blog/models"
	"dev-blog/utils"
	"net/http"
)

const pepper = "iamapepper"

// Form type to represent user sign and login form
type Form struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

func signIn(w http.ResponseWriter, user *models.User) {
	if user.RememberToken == "" {
		token := utils.GenerateRemeberToken()
		user.RememberToken = token
		user.RememberTokenHash = utils.Hash(token)
		db := dbp.New()
		err := db.Save(user).Error
		shouldReturn := utils.MustAndSendError(w, err)
		if shouldReturn {
			return
		}
	}
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    user.RememberToken,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}
