package main

import (
	"fmt"
	"net/http"
	"log"
)

func sendNotification(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Notification sent")
}

func main() {
	http.HandleFunc("/notify", sendNotification)
	log.Println("Notification Service listening on port 8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
