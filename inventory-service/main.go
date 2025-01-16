package main

import (
	"fmt"
	"net/http"
	"log"
)

func checkStock(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Product is in stock")
}

func main() {
	http.HandleFunc("/stock-check", checkStock)
	log.Println("Inventory Service listening on port 8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
