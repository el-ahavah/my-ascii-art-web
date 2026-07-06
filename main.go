package main

import (
	"fmt"
	"net/http"
	"web/handlers"
)

func main() {
	// routes wrapped with middleware
	http.Handle("/", handlers.Logger(http.HandlerFunc(handlers.Home)))
	http.Handle("/ascii-art", handlers.Logger(http.HandlerFunc(handlers.Ascii)))
	http.Handle("/ascii-art-switch", handlers.Logger(http.HandlerFunc(handlers.Switch)))

	// http.HandleFunc("/", handlers.Home)
	// http.HandleFunc("/ascii-art", handlers.Ascii)
	// http.HandleFunc("/ascii-art-switch", handlers.Switch)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
