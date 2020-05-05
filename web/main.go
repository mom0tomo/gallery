package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mom0tomo/chichichimaru_gallery/gallery"
)

func main() {
	http.HandleFunc("/post", garalley.Post)
	http.HandleFunc("/index", garalley.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
