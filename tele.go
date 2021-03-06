package main

import (
	"fmt"
	"log"
//"math/rand"
	"encoding/json"
	"net/http"
	"html"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type AutoGen struct {
	Type  string `json:"type"`
	Value struct {
		      Categories []interface{} `json:"categories"`
		      ID         int           `json:"id"`
		      Joke       string        `json:"joke"`
	      } `json:"value"`
}

func main() {

	//fmt.Printf(data.Value.Joke)

	bot, err := tgbotapi.NewBotAPI("126270440:AAEs6OiKnDpWCL6xsO08T_hvlQKiI6m8tB8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {

		log.Printf(update.Message.CommandArguments())
		log.Printf(update.Message.Command())
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch (update.Message.Text) {
		case "/joke", "/joke@KatKatBot":

			resp, err := http.Get("http://api.icndb.com/jokes/random")

			if err != nil {
				log.Panic(err)
				log.Fatal(err)
			}
			decoder := json.NewDecoder(resp.Body)
			var data AutoGen
			err = decoder.Decode(&data)

			txt := fmt.Sprintf(html.UnescapeString(data.Value.Joke))
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, txt)
			//msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		case "/help", "/help@KatKatBot":
			txt := fmt.Sprintf("Chuck Norris bräuchte keine Hilfe.")
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, txt)
			//msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}

	}
}
