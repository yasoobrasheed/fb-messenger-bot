package send_message

import (
	"bytes"
	"encoding/json"
	"fbmessenger_bot/secrets"
	"fmt"
	"net/http"
)

type Recipient struct {
	ID string `json:"id"`
}

type Payload struct {
	URL        string `json:"url"`
	IsReusable bool   `json:"is_reusable"`
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

func HandleSendMessageText(messageText string, recipientId string) {
	url := buildPageUrl()

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
		fmt.Println("Error marshalling request body:", err)
		return
	}
	requestBody := bytes.NewBuffer(jsonData)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		fmt.Println("Error posting message via Messenger API:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("POST request successful")
	} else {
		fmt.Println("POST request failed with status:", response.Status)
	}
}

func buildPageUrl() string {
	messenger_api_url := "https://graph.facebook.com/v17.0/me/messages"
	access_token_query_param := "?access_token=" + secrets.PAGE_ACCESS_TOKEN
	return messenger_api_url + access_token_query_param
}
