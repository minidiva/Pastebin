package main

import (
	"fmt"
	"log"
	"net/http"
	"pastebin/internal/db"
	"pastebin/internal/handlers"
	"pastebin/internal/repo"
	"pastebin/internal/service"
)

func main() {

	db, err := db.SetupDB()
	if err != nil {
		log.Fatalf("error connect to DB: %v", err)
	}

	r := repo.NewPasteRepo(db)

	s := service.NewPasteService(r)

	PasteHandler := handlers.NewPasteHandler(s)

	http.HandleFunc("/health", PasteHandler.CheckHealth)

	http.HandleFunc("/upload", PasteHandler.CreatePaste)

	fmt.Println("Starting server on localhost:8081...")
	if err = http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("error starting server... %v", err)
	}
}
