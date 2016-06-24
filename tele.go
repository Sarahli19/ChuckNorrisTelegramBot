package main

import (
	"fmt"
	"log"
	"github.com/Syfaro/telegram-bot-api"
	//"math/rand"
	"encoding/json"
	"net/http"

	"html"
)
type AutoGen struct {
	Matches []struct {
		Href string `json:"href"`
		Score int `json:"score"`
		Type string `json:"type"`
		URI string `json:"uri"`
		Value string `json:"value"`
	} `json:"matches"`
	Found int `json:"found"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
}

//423d89be3337fd9ab5fe0abae97a2e45239e3ba4eae8cd1a8aa5
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

			req := http.NewRequest("GET", "http://api.icndb.com/jokes/random", nil)
			req.Header.Set()
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
