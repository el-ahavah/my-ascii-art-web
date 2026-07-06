// package handlers

// import (
// 	"html/template"
// 	"net/http"
// )

// func ErrorHandler(w http.ResponseWriter, file string, status int) {
// 	w.WriteHeader(status)

//		temp, err := template.ParseFiles(file)
//		if err != nil {
//			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//			return
//		}
//		temp.Execute(w, nil)
//	}
package handlers

import (
	"html/template"
	"net/http"
)

type ErrorPage struct {
	Code    int
	Message string
}

func ErrorHandler(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)

	data := ErrorPage{
		Code:    status,
		Message: message,
	}

	temp, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
