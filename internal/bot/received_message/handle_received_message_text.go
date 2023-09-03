package received_message

import (
	"fbmessenger_bot/internal/bot/send_message"
	"fbmessenger_bot/internal/processing"
	"fbmessenger_bot/util"
	"log"
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

func HandleRecievedMessageText(w http.ResponseWriter, r *http.Request) error {
	var messageData MessageData

	err := util.ParseAndUnmarshallRequestBody(r, &messageData)
	if err != nil {
		return err
	} // named return values

	w.WriteHeader(http.StatusOK)

	// Write message to database and ensure only one review is written per order
	messaging := messageData.Entry[0].Messaging[0] // this is poorly written,, must validate first
	userId := messaging.Sender.ID
	if processing.UserReviewExists(messaging.Sender.ID) {
		log.Printf("UserId %s has already written a review for this product and received a discount code.", userId)
		return nil
	} else {
		// userReview := map[string]interface{}{
		// 	"recipient_id": messagingData.Recipient.ID,
		// 	"timestamp":    messagingData.Timestamp,
		// 	"text":         messagingData.Message.Text,
		// }
		userReview := make(map[string]interface{})
		userReview["recipient_id"] = messaging.Recipient.ID
		userReview["timestamp"] = messaging.Timestamp
		userReview["text"] = messaging.Message.Text
		processing.WriteUserReview(userId, userReview)
	}

	// Store review and respond to user
	responseText, err := processing.HandleReceivedReview(messaging.Message.Text)
	if err != nil {
		return err
	}
	err = send_message.HandleSendMessageText(responseText, messaging.Sender.ID)
	if err != nil {
		return err
	}
	return nil
}
