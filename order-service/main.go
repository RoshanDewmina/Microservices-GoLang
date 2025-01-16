package main

import (
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

func placeOrder(w http.ResponseWriter, r *http.Request) {
	authResp, err := http.Get("http://auth-service:8081/verify")
	if err != nil || authResp.StatusCode != http.StatusOK {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	inventoryResp, err := http.Get("http://inventory-service:8085/stock-check")
	if err != nil || inventoryResp.StatusCode != http.StatusOK {
		http.Error(w, "Product not available", http.StatusConflict)
		return
	}

	paymentResp, err := http.Get("http://payment-service:8083/process-payment")
	if err != nil || paymentResp.StatusCode != http.StatusOK {
		http.Error(w, "Payment failed", http.StatusPaymentRequired)
		return
	}

	http.Get("http://notification-service:8084/notify")
	fmt.Fprintln(w, "Order placed successfully")
}

func main() {
	http.HandleFunc("/order", placeOrder)
	log.Println("Order Service listening on port 8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
