package handlers

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, 404, "use the right path")
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, 405, "Method not allowed")
		return
	}

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, 500, "unable to read template")
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, 500, "unable to execute")
	}
}
