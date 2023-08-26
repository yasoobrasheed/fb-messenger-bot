# Facebook Messenger Bot

### Purpose

This bot will send order details to a customer who has received a purchased item. The receipt of the bot will encourage the user to write a review in the messenger window. The user in return will get a discount code for their next purchase or will be routed to a customer service number to assist with any complications.

### Code Walkthrough

Environment Setup
* [Install ngrok](https://ngrok.com/download).
* `ngrok http 8080 --host-header="localhost:8080"`.
* To generate an access token for your messenger app [in the developer portal](https://developers.facebook.com/apps/3611148295794646/messenger/settings/). Paste ngrok url with `/webhook` appended and a verify token, which you will need to input for `secrets.VERIFY_TOKEN` in `messenger_webhook.go`.
* `go run cmd/main.go`.

Send POST Request to http://localhost:8080/order_complete with:
```JSON
{
    "id": "ID",
    "product_name": "Product Name",
	"price": 0.0,
	"delivery_date": "August 26th, 2023",
    "delivery_url": "https://www.deliveryurl.com",
    "sender_id": "Sender ID"
}
```

Now for the user with the sender id from above, you should be able to submit a review to the bot.
