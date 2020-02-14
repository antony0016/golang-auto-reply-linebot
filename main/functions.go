package main

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// HandleAllEvents will handle all events from linebot
func HandleAllEvents(w http.ResponseWriter, req *http.Request) {
	// get events from linebot
	events, err := bot.ParseRequest(req)
	// check linebot have right signature
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400) // 400 Bad Request
		} else {
			w.WriteHeader(500) // Internal Server Error
		}
		return
	}
	// start to process event by message type
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// When server receive a text message
			case *linebot.TextMessage:
				getTextMessage(event, message.Text)
			}
		}
	}
}

func getTextMessage(event *linebot.Event, message string) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do()
	if err != nil {
		log.Print(err)
	}
}
