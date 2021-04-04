package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	stringify "github.com/vicanso/go-stringify"
)

const (
	// APIEndpoint is the endpoint for all API methods,
	// with formatting for Sprintf.
	APIEndpoint = "https://api.telegram.org/bot%s/%s"
	// FileEndpoint is the endpoint for downloading a file from Telegram.
	FileEndpoint = "https://api.telegram.org/file/bot%s/%s"
)

func main() {
	creators := strings.Join(getBotCreators(), ", ")
	log.Printf("GPT3-ru Bot by %s", creators)
	log.Printf("Loading...")
	bot, err := tgbotapi.NewBotAPI(getBotToken())
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	sendMessage(bot, getReportsChat(), "Started successfully!", 0)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		switch update.Message.Command() {
		case "start":
			sendMessage(bot, update.Message.Chat.ID, "Hello. This is a bot based on the GPT3-ru machine learning API by SberCloud.\n\nTo pre-generate text, use the /gen command", 0)
		case "dmsg":
			data := stringify.String(update.Message, replacer)
			sendMessage(bot, update.Message.Chat.ID, data, update.Message.MessageID)
		case "gen":
			go processMessage(bot, update)
		}
	}
}
