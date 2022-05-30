package utils

import (
	"fmt"
	"os"

	"github.com/DisgoOrg/disgohook"
	"github.com/DisgoOrg/disgohook/api"
	"github.com/joho/godotenv"
)

func GetWeebhook() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Hata var", err)
	}
	return os.Getenv("WEBHOOK_URL")
}
func ConnectWebhook() api.WebhookClient {
	webhook, err := disgohook.NewWebhookClientByToken(nil, nil, GetWeebhook())
	if err != nil {
		fmt.Println("Hata vardır", err)
	}
	return webhook
}

var webhook api.WebhookClient = ConnectWebhook()

func SendMessageCreatedAnime(msg string) *api.WebhookMessage {
	message, err := webhook.SendEmbeds(api.NewEmbedBuilder().
		SetTitle("Siteye Yeni Anime Eklendi").
		SetDescription(msg).
		SetURL("https://atlasch.me").
		Build(),
	)
	if err != nil {
		fmt.Println("Hata var", err)
	}
	return message
}
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
