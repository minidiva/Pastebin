package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server on localhost:8080...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("error starting server... %v", err)
	}
}
