package routing

import (
	"net/http"

	"github.com/NullJupiter/GoPresentation/src/routing/routes"
	"github.com/gorilla/mux"
)

// InitRouting function is used to initialize the base routing
func InitRouting() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", routes.IndexGetHandler).Methods(http.MethodGet)

	r.HandleFunc("/presentation/upload", routes.PresentationUploadGetHandler).Methods(http.MethodGet)
	r.HandleFunc("/presentation/upload", routes.PresentationUploadPostHandler).Methods(http.MethodPost)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return r
}
