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
	data := ErrorPage{
		Code:    status,
		Message: message,
	}

	temp, err := template.ParseFiles("templates/erroor.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	err = temp.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// curl -i http://localhost:8080/badroute
