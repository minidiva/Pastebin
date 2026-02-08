package main

import (
	"auth/internal/handler"
	"auth/internal/repo"
	"auth/internal/service"
	"fmt"
	"log"
	"net/http"
)

func main() {

	r := repo.NewUserRepo()

	s := service.NewUserService(r)

	userHandler := handler.NewUserHandler(s)

	http.HandleFunc("/", userHandler.HomeHandler)
	http.HandleFunc("/login", userHandler.LoginHandler)

	fmt.Println("Starting server on localhost:8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error starting server... %v", err)
	}
}
