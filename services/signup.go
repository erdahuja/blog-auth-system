package services

import (
	dbp "dev-blog/db" // rename import of db as dbp
	"dev-blog/models"
	"dev-blog/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Create adds a new user to db
func Create(w http.ResponseWriter, req *http.Request) {
	form := new(Form)
	// get password, email into form struct
	utils.ParseForm(form, req)
	// Add pepper string to password
	pwdWithPepper := form.Password + pepper
	// Create a hash using bcrypt
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(pwdWithPepper), bcrypt.DefaultCost)
	// Check for errors while hashing
	shouldReturn := utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
	// Store in user struct
	user := models.User{
		Email:        form.Email,
		PasswordHash: string(hashedBytes),
	}
	if user.RememberToken != "" {
		user.RememberTokenHash = utils.Hash(user.RememberToken)
	}
	db := dbp.New()
	defer db.Close()
	// Store in database
	err = db.Create(&user).Error
	shouldReturn = utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
	signIn(w, &user)
	http.Redirect(w, req, "/profile", http.StatusFound)
}
