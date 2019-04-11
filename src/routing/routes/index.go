package routes

import (
	"net/http"
)

// IndexGetHandler function is used to handler GET request on /
func IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/presentation/upload", http.StatusSeeOther)
}
