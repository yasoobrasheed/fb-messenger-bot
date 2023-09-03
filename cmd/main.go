package main

import (
	"fbmessenger_bot/internal/bot/received_message"
	"fbmessenger_bot/internal/cars"
	"fbmessenger_bot/internal/processing"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", received_message.HandleWebhook)
	http.HandleFunc("/order_complete", processing.HandleOrderComplete)
	http.HandleFunc("/cars", cars.HandleCars)

	fmt.Println("Server is listening on port 8080...") // switch to log package
	http.ListenAndServe(":8080", nil)                  // this can return an error, catch it
}
