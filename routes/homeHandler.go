package routes

import (
	"dev-blog/utils"
	"net/http"
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	shouldReturn := utils.MustAndSendError(w, err)
	if shouldReturn {
		return
	}
}
