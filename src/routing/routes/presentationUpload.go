package routes

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/NullJupiter/GoPresentation/src/templating"
)

// PresentationUploadGetHandler function is used to handler GET requests from /presentation/upload
func PresentationUploadGetHandler(w http.ResponseWriter, r *http.Request) {
	err := templating.GetTpls().ExecuteTemplate(w, "upload", nil)
	if err != nil {
		log.Fatalf("failed to execute template: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

// PresentationUploadPostHandler function is used to handle POST request from /presentation/upload
func PresentationUploadPostHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, file); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = templating.GetTpls().ExecuteTemplate(w, "presentation", buf)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
