package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"main.go/handle"
)

func main() {
	fmt.Printf("Server run on: http://localhost:8080/\n")
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	http.HandleFunc("/", handle.FormHandler)
	http.HandleFunc("/ascii-art", handle.GenerateHandler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
