package main

import (
	"fmt"
	"g-tracker/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// serve 'static' files
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	// handle 'home' page
	mux.HandleFunc("/", handlers.Home)

	mux.HandleFunc("/artist/", handlers.Artist)

	// handle 'raw' page
	mux.HandleFunc("/raw", handlers.Raw_All)
	mux.HandleFunc("/raw/", handlers.Raw_Artist)

	// start server
	port := "8000"
	fmt.Println("Launching http://localhost:" + port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
