package main

import (
	"challenge-trafilea/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to my API")
	fmt.Println("Server running on the port 3000")
	log.Fatal(http.ListenAndServe(":3000", handlers.RoutesMapper()))
}
