package main

import (
	"log"
	"net/http"

	"github.com/Redict/imgstore/pkg/auth"
)

func main() {
	http.HandleFunc("/register", auth.RegistrationHandler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
