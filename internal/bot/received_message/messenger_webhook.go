package received_message

import (
	"fbmessenger_bot/secrets"
	"log"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	// fb messenger uses GET to establish this as an authorized webhook
	if r.Method == http.MethodGet {
		handleVerification(w, r)
	}

	// fb messenger then uses POST to communicate with this server
	if r.Method == http.MethodPost {
		err := HandleRecievedMessageText(w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func handleVerification(w http.ResponseWriter, r *http.Request) {
	// Our verify token must match fb messenger's verify token
	if secrets.VERIFY_TOKEN != r.URL.Query().Get("hub.verify_token") {
		log.Println("Personal verify token does not match hub verify token.")
		w.WriteHeader(http.StatusUnauthorized)
	}

	// We must return 200 with fb messenger's challenge token
	challengeToken := r.URL.Query().Get("hub.challenge")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(challengeToken))
}
