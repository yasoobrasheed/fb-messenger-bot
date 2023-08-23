package bot

import (
	"fbmessenger_bot/secrets"
	"fmt"
	"io"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	LogAndValidateRequest(w, r)

	if secrets.VERIFY_TOKEN != r.URL.Query().Get("hub.verify_token") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(r.URL.Query().Get("hub.challenge")))
}

func LogAndValidateRequest(w http.ResponseWriter, r *http.Request) {
	// Print the request method, URL, and headers
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
