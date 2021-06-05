package telegram

import (
	"log"
	"io/ioutil"
	"yagan.bot/config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendNotification(message string) {
	bot, err := tgbotapi.NewBotAPI(config.LoadedConfig.Token)
	//bot, err := tgbotapi.NewBotAPI("1785654780:AAEP0ssWGj7Rt0jMxk57PAxbKYSnwvuO9fw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s , sending message ...", bot.Self.UserName)

	msg := tgbotapi.NewMessage(-1001492374088, message)
	bot.Send(msg)
}

func SendImage(path string){
	bot, err := tgbotapi.NewBotAPI(config.LoadedConfig.Token)
	//bot, err := tgbotapi.NewBotAPI("1785654780:AAEP0ssWGj7Rt0jMxk57PAxbKYSnwvuO9fw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s , sending image ...", bot.Self.UserName)

	photoBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photoBytes,
	}

	bot.Send(tgbotapi.NewPhotoUpload(-1001492374088, photoFileBytes))
}
