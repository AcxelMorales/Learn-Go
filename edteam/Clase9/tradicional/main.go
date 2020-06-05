package main

import (
	"log"
	"net/http"
)

func main() {
	// Primera forma de crear un servidor en GO
	// http.Handle("/", http.FileServer(http.Dir("public")))
	// log.Println("Ejecutando servidor en: http://localhost:8080")
	// log.Println(http.ListenAndServe(":8080", nil))

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/", fs)
	log.Println("Ejecutando servidor en: http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
