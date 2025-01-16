package main

import (
	"fmt"
	"net/http"
	"log"
)

func authenticateUser(w http.ResponseWriter, r *http.Request) {
	token := "dummy-token"
	fmt.Fprintln(w, token)
}

func verifyToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "Bearer dummy-token" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/login", authenticateUser)
	http.HandleFunc("/verify", verifyToken)

	log.Println("Auth Service listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
