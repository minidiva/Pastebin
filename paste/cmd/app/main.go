package main

import (
	"fmt"
	"log"
	"net/http"
	"pastebin/internal/handlers"
	"pastebin/internal/repo"
	"pastebin/internal/service"
)

func main() {

	r := repo.NewRepo()

	s := service.NewPasteService(r)

	PasteHandler := handlers.NewPasteHandler(s)

	http.HandleFunc("/health", PasteHandler.CheckHealth)

	fmt.Println("Starting server on localhost:8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("error starting server... %v", err)
	}
}
