package main

import (
	"fmt"
	"github.com/mom0tomo/gallery"
	"log"
	"net/http"
	"os"

	_ "github.com/mom0tomo/gallery"
)

func main() {
	http.HandleFunc("/post", gallery.Post)
	http.HandleFunc("/index", gallery.Index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
