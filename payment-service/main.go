package main

import (
	"fmt"
	"net/http"
	"log"
)

func processPayment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Payment processed")
}

func main() {
	http.HandleFunc("/process-payment", processPayment)
	log.Println("Payment Service listening on port 8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
