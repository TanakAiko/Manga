package handlers

import (
	"net/http"
	"text/template"
)

func RenderHandler(page string) http.HandlerFunc {
	files := []string{"./templates/base.html", "./templates/" + page}
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}
		errr := tmpl.Execute(w, nil)
		if errr != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}
