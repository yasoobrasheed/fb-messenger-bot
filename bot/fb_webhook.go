package bot

import (
	"fbmessenger_bot/secrets"
	"fmt"
	"io"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	// Get insight into fb messenger's payload
	logAndValidateRequest(w, r)

	// fb messenger uses GET to establish this as an authorized webhook
	if r.Method == http.MethodGet {
		handleVerification(w, r)
	}

	// fb messenger then uses POST to communicate with this server
	if r.Method == http.MethodPost {
		HandlePost()
	}
}

func handleVerification(w http.ResponseWriter, r *http.Request) {
	// Our verify token must match fb messenger's verify token
	if secrets.VERIFY_TOKEN != r.URL.Query().Get("hub.verify_token") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(nil)
	}

	// We must return 200 with fb messenger's challenge token
	challengeToken := r.URL.Query().Get("hub.challenge")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(challengeToken))
}

func logAndValidateRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("URL:", r.URL)
	fmt.Println("Headers:", r.Header)
	fmt.Println("Challenge Token:", r.URL.Query().Get("hub.challenge"))
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	fmt.Println("Body:", string(body))
}
