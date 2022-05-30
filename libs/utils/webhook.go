package utils

import (
	"fmt"

	"github.com/Constani/main/libs"
	"github.com/DisgoOrg/disgohook"
	"github.com/DisgoOrg/disgohook/api"
)

// Connected Webhook in disgohook
func ConnectWebhook() api.WebhookClient {
	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, libs.GetWeebhook())
	if err != nil {
		fmt.Println("Hata vardır", err)
	}
	return webhook
}

// Webhook Send Created Anime Message
var webhook api.WebhookClient = ConnectWebhook()

// Send Message example : <utils>.SendMessage(name : "Hello,World", title : "dasdasdads", url: "atlasch.me")
func SendMessage(msg string, title string, url string) *api.WebhookMessage {
	message, err := webhook.SendEmbeds(api.NewEmbedBuilder().
		SetTitle(title).
		SetDescription(msg).
		SetURL(url).
		Build(),
	)
	if err != nil {
		fmt.Println("Hata var", err)
	}
	return message
}
