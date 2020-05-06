package main

import (
	"github.com/Syfaro/telegram-bot-api"
	"log"
)

func main() {

	/* weather, err := owm.NewCurrent("C", "EN", "COPY&PASTE_YOUR_OPENWATHERMAP_TOKEN_HERE")
	if err != nil {
		log.Fatalln(err)
	} */
	// vain attempts to add weather forecasting capabilities to the bot

	// connect to created bot (create a bot at BotFather on Telegram first)
	bot, err := tgbotapi.NewBotAPI("COPY&PASTE_YOUR_TELEGRAM_BOT_TOKEN_HERE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// get updates from the BotFather API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	// read updates
	for {
		select {
		case update := <-upd:
			// User that interacts with the bot
			UserName := update.Message.From.UserName

			// chat ID (ID of the user chat == UserID / ID of a public chat/channel)
			ChatID := update.Message.Chat.ID

			// message text
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// currentWeather := weather.CurrentByName(Text)
			// vain attempts to add weather forecasting capabilities to the bot

			// echo user message
			reply := Text
			// create a message to send to the bot API
			msg := tgbotapi.NewMessage(ChatID, reply)
			// send created message
			bot.Send(msg)
		}

	}


}