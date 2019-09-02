package routes

import (
	"dev-blog/services"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func pageNotFoundFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<strong>Page Not Found</strong>")
}

var tpl *template.Template
var router *mux.Router

func init() {
	tpl = template.Must(template.ParseGlob("views/*"))
	router = mux.NewRouter()
}

// SetupRoutes register routes and handlers
func SetupRoutes() {
	router.HandleFunc("/", homeFunc)
	router.HandleFunc("/login", login).Methods("GET")
	router.HandleFunc("/login", services.Login).Methods("POST")
	router.HandleFunc("/signup", signUp).Methods("GET")
	router.HandleFunc("/signup", services.Create).Methods("POST")
	router.HandleFunc("/profile", profileFunc).Methods("GET")
	router.HandleFunc("/account", accountFunc).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(pageNotFoundFunc)
	http.ListenAndServe(":3000", router)
	fmt.Println("Server listening on PORT :3000")
}
