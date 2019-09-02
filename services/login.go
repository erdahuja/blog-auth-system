package services

import (
	dbp "dev-blog/db"
	"dev-blog/models"
	"dev-blog/utils"
	"net/http"
)

// Login authenticate user with email and password
func Login(w http.ResponseWriter, req *http.Request) {
	form := new(Form)
	utils.ParseForm(form, req)
	db := dbp.New()
	var user models.User
	err := db.Where("email=?", form.Email).First(&user).Error
	shouldReturn := utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
	err = utils.CompareHashAndPassword([]byte(user.PasswordHash), []byte(form.Password+pepper))
	shouldReturn = utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
	signIn(w, &user)
	db.Close()
	http.Redirect(w, req, "/profile", http.StatusFound)
}
