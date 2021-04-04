package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func processMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	args := strings.Split(update.Message.Text, " ")
	if len(args) == 1 {
		sendMessage(bot, update.Message.Chat.ID, "Invalid syntax! Example: /gen Society if Kuba stfu:\n1) ", update.Message.MessageID)
		return
	}
	args = RemoveIndex(args, 0)
	text := strings.Join(args, " ")
	message, err := sendMessage(bot, update.Message.Chat.ID, "Generating...", update.Message.MessageID)
	if err != nil {
		log.Panic(err)
	}
	request, err := makeRequest(text)
	if err != nil {
		editMessage(bot, update.Message.Chat.ID, message.MessageID, "An error occured: "+err.Error())
		return
	}
	answer := parseAPIAnswer([]byte(request))
	if answer.Detail != nil {
		editMessage(bot, update.Message.Chat.ID, message.MessageID, "An error occured: "+answer.Detail[0].Message)
		return
	}
	editMessage(bot, update.Message.Chat.ID, message.MessageID, answer.Predictions)
}
