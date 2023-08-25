package received_message

import (
	"fbmessenger_bot/internal/bot/send_message"
	"fbmessenger_bot/internal/processing"
	"fbmessenger_bot/util"
	"fmt"
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
	var messageData MessageData

	err := util.ParseAndUnmarshallRequestBody(r, &messageData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Write message to database and ensure only one review is written per order
	messaging := messageData.Entry[0].Messaging[0]
	userId := messaging.Sender.ID
	if processing.UserReviewExists(messaging.Sender.ID) {
		// TODO: make this some kind of backup message (or none at all)
		return
	} else {
		userReview := make(map[string]interface{})
		userReview["recipient_id"] = messaging.Recipient.ID
		userReview["timestamp"] = messaging.Timestamp
		userReview["text"] = messaging.Message.Text
		processing.WriteUserReview(userId, userReview)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received successfully"))

	messageText := processing.HandleReceivedReview(messaging.Message.Text)
	send_message.HandleSendMessageText(messageText, messaging.Sender.ID)
}
