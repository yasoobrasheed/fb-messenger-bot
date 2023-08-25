package received_message

import (
	"encoding/json"
	"fbmessenger_bot/internal/bot/send_message"
	"fbmessenger_bot/internal/processing"
	"fmt"
	"io"
	"net/http"
)

type MessageData struct {
	Entry  []Entry `json:"entry"`
	Object string  `json:"object"`
}

type Entry struct {
	ID        string      `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

type Messaging struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   Message   `json:"message"`
}

type Sender struct {
	ID string `json:"id"`
}

type Recipient struct {
	ID string `json:"id"`
}

type Message struct {
	Mid  string `json:"mid"`
	Text string `json:"text"`
}

func HandleRecievedMessageText(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var messageData MessageData

	err = json.Unmarshal([]byte(requestBody), &messageData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	messaging := messageData.Entry[0].Messaging[0]

	// TODO: Write review to DB if this is the first response from the user

	// fmt.Println("Sender ID:", messaging.Sender.ID)
	// fmt.Println("Recipient ID:", messaging.Recipient.ID)
	// fmt.Println("Timestamp:", messaging.Timestamp)
	// fmt.Println("Text:", messaging.Message.Text)

	// TODO: If not first response from user, send generic message to check site FAQ

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received successfully"))

	messageText := processing.HandleReceivedReview(messaging.Message.Text)
	send_message.HandleSendMessageText(messageText, messaging.Sender.ID)
}
