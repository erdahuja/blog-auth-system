package routes

import (
	"dev-blog/utils"
	"net/http"
)

func signUp(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	shouldReturn := utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
}
