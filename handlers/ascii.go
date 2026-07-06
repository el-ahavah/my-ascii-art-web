package handlers

import (
	"html/template"
	"net/http"
	"web/ascii_art"
)

type PageData struct {
	Result string
	Text   string
	Banner string
}

var temp = template.Must(template.ParseFiles("templates/index.html"))

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		ErrorHandler(w, 404, "use the right path")
		return
	}

	if r.Method != http.MethodPost {
		ErrorHandler(w, 405, "Method not allowed")
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")

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
		ErrorHandler(w, 500, "unable to generate ascii")
		return
	}

	data := PageData{
		Result: result,
		Text:   text,
		Banner: banner,
	}

	err = temp.Execute(w, data)
	if err != nil {
		ErrorHandler(w, 500, "unable to execute ascii")
		return
	}
}
