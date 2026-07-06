package handlers

import (
	"html/template"
	"net/http"
	"web/ascii_art"
)

func Switch(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art-switch" {
		ErrorHandler(w, 404, "use the right path")
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, 405, "Method not allowed")
		return
	}

	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		ErrorHandler(w, 404, "invalid banner")
		return
	}

	if text == "" || banner == "" {
		ErrorHandler(w, 400, "make provisions for text and banner")
		return
	}

	autocomplete := "banners/" + banner + ".txt"
	result, err := ascii_art.GenerateAscii(text, autocomplete)
	if err != nil {
		ErrorHandler(w, 500, "unable to generate ascii for switch")
		return
	}

	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, 500, "unable to read template")
		return
	}

	data := PageData{
		Result: result,
		Text:   text,
		Banner: banner,
	}

	err = temp.Execute(w, data)
	if err != nil {
		ErrorHandler(w, 500, "unable to execute ascii for switch")
		return
	}
}
