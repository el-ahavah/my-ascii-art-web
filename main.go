package main

import (
	"fmt"
	"my-ascii-art-web/handlers"
	"net/http"
)

func main() {
	// Configure a file server rooted at the `static` directory
	// Mount it under the `/static/` URL path.
	fs := http.FileServer(http.Dir("static"))
	
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.Handle("/static/", handlers.Logger(http.StripPrefix("/static/", fs),),)

	http.Handle("/static/", http.StripPrefix("/static/", handlers.Logger(fs),),)

	// routes wrapped with middleware
	http.Handle("/", handlers.Logger(http.HandlerFunc(handlers.Home)))
	http.Handle("/ascii-art", handlers.Logger(http.HandlerFunc(handlers.Ascii)))
	http.Handle("/ascii-art-switch", handlers.Logger(http.HandlerFunc(handlers.Switch)))

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
