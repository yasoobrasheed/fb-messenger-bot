package main

import (
	"fbmessenger_bot/bot"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", bot.HandleWebhook)

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
