package routes

import (
	"dev-blog/services"
	"dev-blog/utils"
	"net/http"
)

func profileFunc(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	shouldReturn := utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
	user, err := services.ByRemember(cookie.Value)
	shouldReturn = utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
	err = tpl.ExecuteTemplate(w, "profile.gohtml", user)
	shouldReturn = utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
}
