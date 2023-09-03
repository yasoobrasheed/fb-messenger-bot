package send_message

import (
	"bytes"
	"encoding/json"
	"fbmessenger_bot/secrets"
	"log"
	"net/http"
)

type Recipient struct {
	ID string `json:"id"`
}

type Message struct {
	Text string `json:"text"`
}

type RequestBody struct {
	Recipient     Recipient `json:"recipient"`
	Message       Message   `json:"message"`
	MessagingType string    `json:"messaging_type"`
	Tag           string    `json:"tag"`
}

func HandleSendMessageText(messageText string, recipientId string) error {
	data := RequestBody{
		Recipient:     Recipient{ID: recipientId},
		MessagingType: "MESSAGE_TAG",
		Message: Message{
			Text: messageText,
		},
		Tag: "CONFIRMED_EVENT_UPDATE",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = makeGraphAPIRequest(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func makeGraphAPIRequest(jsonData []byte) (*http.Response, error) {
	graphAPIUrl := buildGraphAPIUrl()

	requestBody := bytes.NewBuffer(jsonData)

	response, err := http.Post(graphAPIUrl, "application/json", requestBody)
	if err != nil {
		log.Println("Error posting message via FB Graph API:", err)
		return nil, err
	}
	defer response.Body.Close()

	return response, nil
}

func buildGraphAPIUrl() string {
	messenger_api_url := "https://graph.facebook.com/v17.0/me/messages"
	access_token_query_param := "?access_token=" + secrets.PAGE_ACCESS_TOKEN
	return messenger_api_url + access_token_query_param
}
