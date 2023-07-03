package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"main.go/handle"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4040" // Default port if PORT environment variable is not set
	}

	fmt.Printf("Server run on: http://localhost:%s/\n", PORT)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	http.HandleFunc("/", handle.FormHandler)
	http.HandleFunc("/ascii-art", handle.GenerateHandler)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
