package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/makhmudvazeez/go-postgres/router"
)

func main() {
	r := router.Router()

	fmt.Println("Server starting on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
