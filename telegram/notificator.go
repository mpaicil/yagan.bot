package telegram

import (
	"log"
	"yagan.bot/config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendNotification(message string) {
	bot, err := tgbotapi.NewBotAPI(config.LoadedConfig.Token)
	//bot, err := tgbotapi.NewBotAPI("1785654780:AAEP0ssWGj7Rt0jMxk57PAxbKYSnwvuO9fw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	msg := tgbotapi.NewMessage(-1001492374088, message)
	bot.Send(msg)

}
